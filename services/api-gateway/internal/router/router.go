package router

import (
	"github.com/gofiber/fiber/v2"
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
		user.Get("/", s.ListUsers)
	})

	v1.Route("/companies", func(comp fiber.Router) {
		comp.Use(s.RequireAuthService())
		comp.Use(s.Authorize())

		comp.Get("/", s.ListCompanies)
		comp.Get("/:id", s.GetCompany)
	}, "companies")

	v1.Route("/docs", func(docs fiber.Router) {
		docs.Post("/", s.CreateReportTemplate)
		docs.Get("/:docID", s.GetDocument)
		docs.Get("/", s.ListDocuments)
		docs.Put("/:docID", s.UpdateDocument)
		docs.Get(":docID/download", s.DownloadDocument)
	})

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
