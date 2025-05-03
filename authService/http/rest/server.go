package rest

import (
	"authService/configs"
	"authService/http/rest/handlers"
	"authService/internal/servicesStorage"
	"authService/pkg/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config configs.Config
	gin    *gin.Engine
}

func NewServer() (*Server, error) {
	cnf, err := configs.NewParsedConfig()
	if err != nil {
		return nil, err
	}

	database, err := db.Connect(cnf.Database)

	if err != nil {
		return nil, err
	}

	servicesStore, err := servicesStorage.NewServicesStorage(cnf, database)
	if err != nil {
		return nil, err
	}

	engine := gin.Default()
	handlers.Register(engine, servicesStore)

	s := Server{
		config: cnf,
		gin:    engine,
	}

	return &s, nil
}

func (s *Server) Run() error {
	err := s.gin.Run(s.config.ServerPort)

	if err != nil {
		return err
	}

	return nil
}
