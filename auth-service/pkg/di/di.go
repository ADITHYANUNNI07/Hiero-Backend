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
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	employerRepository := repository.NewEmployerRepository(gormDB)
	employerUseCase := usecase.NewEmployerUseCase(employerRepository)
	employerServiceServer := service.NewEmployerServer(employerUseCase)

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminServiceServer := service.NewAdminServer(adminUseCase)

	jobSeekerRepository := repository.NewJobSeekerRepository(gormDB)
	jobSeekerUseCase := usecase.NewJobSeekerUseCase(jobSeekerRepository)
	jobSeekerServiceServer := service.NewJobSeekerServer(jobSeekerUseCase)

	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer, employerServiceServer, jobSeekerServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
