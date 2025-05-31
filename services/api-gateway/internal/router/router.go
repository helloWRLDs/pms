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
		comp.Get("/:companyID", s.GetCompany)
		comp.Post("/", s.CreateCompany)
		comp.Get("/:companyID/stats", s.GetCompanyStats)

		comp.Route("/:companyID/participants", func(participants fiber.Router) {
			participants.Use(s.Authorize(), s.RequireCompany())

			participants.Post("/", s.CompanyAddParticipant)
			participants.Delete("/:userID", s.CompanyRemoveParticipant)
		})
	})

	v1.Route("/docs", func(docs fiber.Router) {
		docs.Post("/", s.CreateReportTemplate)
		docs.Get("/:docID", s.GetDocument)
		docs.Get("/", s.ListDocuments)
		docs.Put("/:docID", s.UpdateDocument)
		docs.Get(":docID/download", s.DownloadDocument)
	})

	v1.Route("/tasks", func(tasks fiber.Router) {
		tasks.Use(s.RequireAuthService(), s.RequireProjectService())
		tasks.Use(s.Authorize())
		tasks.Use(s.RequireCompany(), s.RequireProject()) // X-Company-ID, X-Project-ID

		tasks.Post("/", s.CreateTask)
		tasks.Get("/", s.ListTasks)
		tasks.Get("/:taskID", s.GetTask)
		tasks.Put("/:taskID", s.UpdateTask)
		tasks.Delete("/:taskID", s.DeleteTask)

		tasks.Route("/:taskID/assignments", func(assignment fiber.Router) {
			assignment.Post("/:userID", s.CreateTaskAssignment)
			assignment.Delete("/:userID", s.DeleteTaskAssignment)
		})

		tasks.Route("/:taskID/comments", func(comment fiber.Router) {
			comment.Get("/", s.ListTaskComments)
			comment.Post("/", s.CreateTaskComments)
		})
	})

	v1.Route("/sprints", func(sprints fiber.Router) {
		sprints.Use(s.Authorize())
		sprints.Use(s.RequireProjectService())
		sprints.Use(s.RequireCompany(), s.RequireProject())

		sprints.Post("/", s.CreateSprint)
		sprints.Get("/", s.ListSprints)
		sprints.Get("/:sprintID", s.GetSprint)
		sprints.Put("/:sprintID", s.UpdateSprint)
	})

	v1.Route("/projects", func(proj fiber.Router) {
		proj.Use(s.Authorize())
		proj.Use(s.RequireAuthService(), s.RequireProjectService())
		proj.Use(s.RequireCompany())

		proj.Post("/", s.CreateProject)
		proj.Get("/", s.ListProjects)
		proj.Get("/:projectID", s.GetProject)
	})

	v1.Route("/background-tasks", func(tasks fiber.Router) {
		tasks.Get("/", s.ListBackgroundTasks)
	})
}
