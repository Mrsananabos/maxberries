package rest

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/http/rest/handlers"
	services "backgroundWorkerService/internal/servicesStorage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config configs.Config
	gin    *gin.Engine
}

func NewServer(cnf configs.Config, services services.ServicesStorage) (*Server, error) {
	engine := gin.Default()
	handlers.Register(engine, services)

	s := Server{
		config: cnf,
		gin:    engine,
	}

	return &s, nil
}

func (s *Server) Run() error {
	err := s.gin.Run(s.config.Port)

	if err != nil {
		return err
	}

	return nil
}
