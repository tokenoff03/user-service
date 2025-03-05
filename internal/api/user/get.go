package user

import (
	"context"
	"user-service/internal/converter"
	"user-service/pkg/user_v1"

	"github.com/pkg/errors"
)

func (i *Implementation) Get(ctx context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	if req.GetId() == 0 {
		return nil, errors.Errorf("id is empty")
	}

	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &user_v1.GetResponse{
		User: converter.ToProtoUser(user),
	}, nil
}
