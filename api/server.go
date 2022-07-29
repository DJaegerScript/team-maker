package api

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Server struct {
	Router *fiber.App
	DB     *pgxpool.Pool
}

func MakeServer(db *pgxpool.Pool) Server {
	r := fiber.New()
	server := Server{
		Router: r,
		DB:     db,
	}
	return server
}

func (s *Server) RunServer() {
	//s.SetupSwagger()
	s.SetupRouter()

	port := os.Getenv("APP_PORT")
	err := s.Router.Listen(":" + port)
	if err != nil {
		zap.L().Fatal("Failed to listen port "+port, zap.Error(err))
	}
}

//func (s *Server) SetupSwagger() {
//
//	docs.SwaggerInfo.Title = "Swagger API"
//	docs.SwaggerInfo.Description = "Cards"
//	docs.SwaggerInfo.Version = "1.0"
//	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
//	docs.SwaggerInfo.Schemes = []string{"http"}
//
//	s.Router.Get("/swagger/*", fiberSwagger.WrapHandler)
//}
