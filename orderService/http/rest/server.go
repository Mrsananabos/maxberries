package rest

import (
	"github.com/gin-gonic/gin"
	"orderService/configs"
	"orderService/http/rest/handlers"
	services "orderService/internal/servicesStorage"
	"orderService/pkg/db"
	"orderService/pkg/kafka"
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

	kafkaProducer, err := kafka.CreateProducer(cnf.Kafka)
	if err != nil {
		return nil, err
	}

	engine := gin.Default()

	serviceStorage := services.NewServicesStorage(cnf, database, kafkaProducer)
	handlers.Register(engine, serviceStorage)

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
