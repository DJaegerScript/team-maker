package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"team-maker/pkg/team"
)

func (s *Server) SetupRouter() {
	s.Router.Use(recover.New())

	api := s.Router.Group("/api")

	api.Mount("/team", s.teamRouter())
}

func (s *Server) teamRouter() *fiber.App {
	teamRepo := team.NewRepo(s.DB)
	teamService := team.NewService(teamRepo)
	teamHandler := team.NewHandler(teamService)

	s.Router.Get("/", teamHandler.GetAllTeams)
	s.Router.Get("/:id", teamHandler.GetTeam)
	s.Router.Post("/", teamHandler.CreateTeam)
	s.Router.Put("/:id", teamHandler.UpdateTeam)
	s.Router.Delete("/:id", teamHandler.DeleteTeam)

	return s.Router
}
