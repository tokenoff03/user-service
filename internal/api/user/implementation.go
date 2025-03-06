package user

import (
	"context"
	"user-service/internal/service"
	"user-service/pkg/user_v1"

	"github.com/tokenoff03/authentication-service/pkg/auth_v1"
	"github.com/tokenoff03/lib_ad1lek/pkg/sys/validate"
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

func validateID(id int64) validate.Condition {
	return func(ctx context.Context) error {
		if id <= 0 {
			return validate.NewValidationErrors("id must be greater than 0")
		}

		return nil
	}
}

func otherValidateID(id int64) validate.Condition {
	return func(ctx context.Context) error {
		if id <= 100 {
			return validate.NewValidationErrors("id must be greater than 100")
		}

		return nil
	}
}
