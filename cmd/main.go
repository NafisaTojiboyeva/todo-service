package main

import (
	"net"

	"app/config"
	pb "app/genproto"
	"app/pkg/db"
	"app/pkg/logger"
	"app/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "todo-service")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	taskService := service.NewToDoService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterToDoServiceServer(s, taskService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
