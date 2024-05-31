package interfaces

import "chat/pkg/utils"

type NewauthClient interface {
	CheckUserAvalilabilityWithUserID(userID int) bool
	UserData(userID int) (models.UserData, error)
}
