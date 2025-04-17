package rest

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/http/rest/handlers"
	"backgroundWorkerService/pkg/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config configs.Config
	gin    *gin.Engine
}

func NewServer() (*Server, error) {
	cnf := configs.NewParsedConfig()

	redis, err := db.Connect(cnf.Redis)
	if err != nil {
		return &Server{}, err
	}

	engine := gin.Default()
	handlers.Register(engine, cnf, redis)

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
