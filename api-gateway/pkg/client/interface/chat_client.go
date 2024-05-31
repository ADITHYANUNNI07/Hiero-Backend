package interfaces

import "HireoGateWay/pkg/utils/models"

type ChatClient interface {
	GetChat(userID string, req models.ChatRequest) ([]models.TempMessage, error)
}
