package router

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"pms.pkg/transport/grpc/dto"
)

func (s *Server) SetupREST() {
	api := s.Group("/api")

	s.Use(s.SecureHeaders())

	v1 := api.Group("/v1")

	v1.Get("/healthcheck", s.HealthcheckHandler)

	v1.Route("/auth", func(auth fiber.Router) {
		auth.Use(s.RequireAuthService())

		auth.Post("/login", s.LoginUser)
		auth.Post("/register", s.RegisterUser)
	})

	v1.Route("/session", func(session fiber.Router) {
		session.Use(s.Authorize())

		session.Get("/", s.GetSession)
		session.Put("/", s.UpdateSession)
		session.Delete("/", s.DeleteSession)
	})

	v1.Route("/users", func(user fiber.Router) {
		user.Use(s.RequireAuthService())
		user.Use(s.Authorize())

		user.Get("/:id", s.GetUser)
	})

	v1.Route("/companies", func(comp fiber.Router) {
		comp.Use(s.RequireAuthService())
		comp.Use(s.Authorize())

		comp.Get("/", s.ListCompanies)
		comp.Get("/:id", s.GetCompany)
	}, "companies")

	v1.Route("/projects", func(proj fiber.Router) {
		proj.Use(s.RequireAuthService(), s.Authorize())

		proj.Get("/", s.ListProjects) // /projects?company_id required
		proj.Get("/:projectID", s.GetProject)
		proj.Post("/", s.CreateProject)

		proj.Route("/:projectID/tasks", func(tasks fiber.Router) {
			tasks.Use(s.CheckCompany())

			tasks.Post("/", s.CreateTask)
			tasks.Get("/", s.ListTasks)
			tasks.Get("/:taskID", s.GetTask)
			tasks.Put("/:taskID", s.UpdateTask)
			tasks.Delete("/:taskID", s.DeleteTask)

			tasks.Route("/:taskID/comments", func(comment fiber.Router) {
				comment.Get("/", s.ListTaskComments)
				comment.Post("/", s.CreateTaskComments)
				comment.Get("/:commentID", s.GetTaskComment)
			})
		})

		proj.Route("/:projectID/sprints", func(sprints fiber.Router) {
			sprints.Use(s.CheckCompany())

			sprints.Post("/", s.CreateSprint)
			sprints.Get("/", s.ListSprints)
			sprints.Get("/:sprintID", s.GetSprint)
			sprints.Put("/:sprintID", s.UpdateSprint)
		})
	})

	v1.Route("/background-tasks", func(tasks fiber.Router) {
		tasks.Get("/", s.ListBackgroundTasks)
	})
}

func (s *Server) SetupWS() {
	s.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)

			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	s.Get("/ws/projects/:projectID/sprints/:sprintID", websocket.New(s.StreamSprint))

	s.Get("/ws/dashboard/:id", websocket.New(s.DashboardStream))

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if s.Logic.ProjectClient() == nil {
				log.Infow("project service unavailable yet")
				continue
			}
			for hubID, hub := range s.wshubs {
				parts := strings.SplitN(hubID, "-", 2)
				if len(parts) != 2 {
					log.Warnw("invalid hub ID", "id", hubID)
					continue
				}
				ent, id := parts[0], parts[1]
				log.Infow("ws hub info", "type", ent, "id", id)

				switch ent {
				case "sprint":
					func() (err error) {
						task, err := s.Logic.TaskQueue.Rpop(context.Background(), hubID)
						if err != nil || task.Value == nil {
							log.Infow("no tasks found in queue")
							return
						}
						if err := s.Logic.UpdateTask(context.Background(), task.Value.Id, task.Value); err != nil {
							log.Errorw("failed to update task", "err", err)
							return err
						}
						log.Infow("task is updated")
						tasks, err := s.Logic.ListTasks(context.Background(), &dto.TaskFilter{
							SprintId: id,
							// AssigneeId: c.Query("assignee_id"),
							Page:    1,
							PerPage: 10000,
						})
						if err != nil {
							log.Errorw("failed to fetch tasks", "err", err)
							return err
						}
						msg, err := json.Marshal(tasks.Items)
						if err != nil {
							log.Errorw("failed marshaling trask list", "err", err)
							return
						}
						log.Info("broadcasting to all connected clients")
						hub.Broadcast(msg)
						if len(hub.GetClients()) == 0 {
							log.Infow("no more clients. removing from queue", "hub_id", id)
							delete(s.wshubs, id)
						}
						return nil
					}()
				}
			}
		}
	}()

	go func() { // dashboard chages db writer
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if len(s.DashboardHub.GetClients()) == 0 {
				log.Infow("no clients connected")
				continue
			}
			if len(s.DashboardHub.GetCache()) == 0 {
				log.Infow("no tasks to save")
				continue
			}
			log.Infow("saving tasks to db...", "tasks", s.DashboardHub.GetCache())
			s.DashboardHub.Clean()
		}
	}()
}
