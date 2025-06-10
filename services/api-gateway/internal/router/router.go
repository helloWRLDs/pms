package router

import (
	"github.com/gofiber/fiber/v2"
	"pms.pkg/consts"
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
		auth.Get("/oauth2/:provider", s.InitiateOAuth2)
		auth.Get("/oauth2/:provider/callback", s.OAuth2Callback)
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
		user.Put("/:id", s.RequireCompany(), s.RequirePermission(consts.USER_WRITE_PERMISSION), s.UpdateUser)
	})

	v1.Route("/companies", func(comp fiber.Router) {
		comp.Use(s.RequireAuthService())
		comp.Use(s.Authorize())

		comp.Get("/", s.ListCompanies)
		comp.Get("/:companyID", s.GetCompany)
		comp.Get("/:companyID/stats", s.GetCompanyStats)

		comp.Post("/", s.CreateCompany)

		comp.Route("/:companyID/participants", func(participants fiber.Router) {
			participants.Use(s.RequireCompanyFromPath())
			participants.Post("/:userID", s.RequirePermission(consts.COMPANY_INVITE_PERMISSION), s.CompanyAddParticipant)
			participants.Delete("/:userID", s.RequirePermission(consts.USER_DELETE_PERMISSION), s.CompanyRemoveParticipant)
		})
	})

	v1.Route("/docs", func(docs fiber.Router) {
		docs.Use(s.Authorize())
		docs.Use(s.RequireCompany())

		docs.Post("/", s.RequirePermission(consts.COMPANY_WRITE_PERMISSION), s.CreateReportTemplate)
		docs.Get("/:docID", s.RequirePermission(consts.COMPANY_READ_PERMISSION), s.GetDocument)
		docs.Get("/", s.RequirePermission(consts.COMPANY_READ_PERMISSION), s.ListDocuments)
		docs.Put("/:docID", s.RequirePermission(consts.COMPANY_WRITE_PERMISSION), s.UpdateDocument)
		docs.Get(":docID/download", s.RequirePermission(consts.COMPANY_READ_PERMISSION), s.DownloadDocument)
	})

	v1.Route("/tasks", func(tasks fiber.Router) {
		tasks.Use(s.RequireAuthService(), s.RequireProjectService())
		tasks.Use(s.Authorize())
		tasks.Use(s.RequireCompany(), s.RequireProject())

		tasks.Post("/", s.RequirePermission(consts.TASK_WRITE_PERMISSION), s.CreateTask)
		tasks.Get("/", s.RequirePermission(consts.TASK_READ_PERMISSION), s.ListTasks)
		tasks.Get("/:taskID", s.RequirePermission(consts.TASK_READ_PERMISSION), s.GetTask)
		tasks.Put("/:taskID", s.RequirePermission(consts.TASK_WRITE_PERMISSION), s.UpdateTask)
		tasks.Delete("/:taskID", s.RequirePermission(consts.TASK_DELETE_PERMISSION), s.DeleteTask)

		tasks.Route("/:taskID/assignments", func(assignment fiber.Router) {
			assignment.Post("/:userID", s.RequirePermission(consts.TASK_INVITE_PERMISSION), s.CreateTaskAssignment)
			assignment.Delete("/:userID", s.RequirePermission(consts.TASK_DELETE_PERMISSION), s.DeleteTaskAssignment)
		})

		tasks.Route("/:taskID/comments", func(comment fiber.Router) {
			comment.Get("/", s.RequirePermission(consts.TASK_READ_PERMISSION), s.ListTaskComments)
			comment.Post("/", s.RequirePermission(consts.TASK_WRITE_PERMISSION), s.CreateTaskComments)
		})
	})

	v1.Route("/roles", func(roles fiber.Router) {
		roles.Use(s.RequireAuthService())
		roles.Use(s.Authorize())
		roles.Get("/", s.ListRoles)

		roles.Use(s.RequireCompany())
		roles.Post("/", s.RequirePermission(consts.ROLE_WRITE_PERMISSION), s.CreateRole)
		roles.Get("/:roleName", s.RequirePermission(consts.ROLE_READ_PERMISSION), s.GetRole)
		roles.Put("/:roleName", s.RequirePermission(consts.ROLE_WRITE_PERMISSION), s.UpdateRole)
		roles.Delete("/:roleName", s.RequirePermission(consts.ROLE_DELETE_PERMISSION), s.DeleteRole)
	})

	v1.Route("/sprints", func(sprints fiber.Router) {
		sprints.Use(s.Authorize())
		sprints.Use(s.RequireProjectService())
		sprints.Use(s.RequireCompany(), s.RequireProject())

		sprints.Post("/", s.RequirePermission(consts.SPRINT_WRITE_PERMISSION), s.CreateSprint)
		sprints.Get("/", s.RequirePermission(consts.SPRINT_READ_PERMISSION), s.ListSprints)
		sprints.Get("/:sprintID", s.RequirePermission(consts.SPRINT_READ_PERMISSION), s.GetSprint)
		sprints.Put("/:sprintID", s.RequirePermission(consts.SPRINT_WRITE_PERMISSION), s.UpdateSprint)
	})

	v1.Route("/projects", func(proj fiber.Router) {
		proj.Use(s.Authorize())
		proj.Use(s.RequireAuthService(), s.RequireProjectService())
		proj.Use(s.RequireCompany())

		proj.Post("/", s.RequirePermission(consts.PROJECT_WRITE_PERMISSION), s.CreateProject)
		proj.Get("/", s.RequirePermission(consts.PROJECT_READ_PERMISSION), s.ListProjects)
		proj.Get("/:projectID", s.RequirePermission(consts.PROJECT_READ_PERMISSION), s.GetProject)
	})

	v1.Route("/background-tasks", func(tasks fiber.Router) {
		tasks.Use(s.Authorize())
		tasks.Use(s.RequireCompany())

		tasks.Get("/", s.RequirePermission(consts.COMPANY_READ_PERMISSION), s.ListBackgroundTasks)
	})
}
