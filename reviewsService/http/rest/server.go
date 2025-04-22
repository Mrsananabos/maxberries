package rest

import (
	"github.com/gin-gonic/gin"
	"reviewsService/configs"
	"reviewsService/http/rest/handlers"
	"reviewsService/pkg/db/mongo"
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

	database, err := mongo.ConnectToCollection(cnf.Mongo)
	if err != nil {
		return nil, err
	}

	engine := gin.Default()
	handlers.Register(engine, database, cnf.Services)

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
