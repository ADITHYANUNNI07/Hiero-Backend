package di

import (
	server "Auth/pkg/api"
	"Auth/pkg/api/service"
	"Auth/pkg/config"
	"Auth/pkg/db"
	"Auth/pkg/repository"
	"Auth/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	// Connect to the database
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	jobRepository := repository.NewJobRepository(gormDB)
	jobUseCase := usecase.NewJobUseCase(jobRepository)
	jobServiceServer := service.NewJobServer(jobUseCase)

	grpcServer, err := server.NewGRPCServer(cfg, jobServiceServer)
	if err != nil {
		return nil, err
	}

	return grpcServer, nil
}
