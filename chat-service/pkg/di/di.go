package di

import (
	server "chat/pkg/api"
	"chat/pkg/api/service"
	"chat/pkg/client"
	"chat/pkg/config"
	"chat/pkg/db"
	"chat/pkg/repository"
	"chat/pkg/usecase"
	"log"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	database, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	chatRepository := repository.NewChatRepository(database)
	authClient := client.NewAuthClient(&cfg)

	chatUseCase := usecase.NewChatUseCase(chatRepository, authClient.Client)

	serviceServer := service.NewChatServer(chatUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, serviceServer)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to MongoDB at", cfg.DBUri)
	log.Println("Connected to database:", database.Name())

	go chatUseCase.MessageConsumer()
	return grpcServer, nil
}
