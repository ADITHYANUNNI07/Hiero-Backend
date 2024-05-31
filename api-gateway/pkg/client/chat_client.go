package client

import (
	"HireoGateWay/pkg/config"
	pb "HireoGateWay/pkg/pb/chat"
	"HireoGateWay/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type ChatClient struct {
	Client pb.ChatServiceClient
}

func NewChatClient(cfg config.Config) *ChatClient {
	grpcConnection, err := grpc.Dial(cfg.ChatSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewChatServiceClient(grpcConnection)

	return &ChatClient{
		Client: grpcClient,
	}
}

func (c *ChatClient) GetChat(userID string, req models.ChatRequest) ([]models.TempMessage, error) {
	if c.Client == nil {
		return nil, fmt.Errorf("gRPC client is not initialized")
	}

	data, err := c.Client.GetFriendChat(context.Background(), &pb.GetFriendChatRequest{
		UserID:   userID,
		FriendID: req.FriendID,
		OffSet:   req.Offset,
		Limit:    req.Limit,
	})
	if err != nil {
		return nil, err
	}

	var response []models.TempMessage
	for _, v := range data.FriendChat {
		chatResponse := models.TempMessage{
			SenderID:    v.SenderId,
			RecipientID: v.RecipientId,
			Content:     v.Content,
			Timestamp:   v.Timestamp,
		}
		response = append(response, chatResponse)
	}

	return response, nil
}
