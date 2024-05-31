package di

import (
	server "HireoGateWay/pkg/api"
	"HireoGateWay/pkg/api/handler"
	"HireoGateWay/pkg/client"
	"HireoGateWay/pkg/config"
	"HireoGateWay/pkg/helper"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	employerClient := client.NewEmployerClient(cfg)
	employerHandler := handler.NewEmployerHandler(employerClient)

	jobSeekerClient := client.NewJobSeekerClient(cfg)
	jobSeekerHandler := handler.NewJobSeekerHandler(jobSeekerClient)

	jobClient := client.NewJobClient(cfg)
	jobHandler := handler.NewJobHandler(jobClient)

	helper := helper.NewHelper(&cfg)
	chatClient := client.NewChatClient(cfg)
	chatHandler := handler.NewChatHandler(chatClient, helper)

	serverHTTP := server.NewServerHTTP(adminHandler, employerHandler, jobSeekerHandler, jobHandler, chatHandler)

	return serverHTTP, nil
}
