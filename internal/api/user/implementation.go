package user

import (
	"user-service/internal/service"
	"user-service/pkg/user_v1"

	"github.com/tokenoff03/authentication-service/pkg/auth_v1"
)

type Implementation struct {
	user_v1.UnimplementedUserV1Server
	userService       service.UserService
	authServiceClient auth_v1.AuthV1Client
}

func NewImplementation(userService service.UserService, authServiceClient auth_v1.AuthV1Client) *Implementation {
	return &Implementation{
		userService:       userService,
		authServiceClient: authServiceClient,
	}
}
