package user

import (
	"context"
	"user-service/internal/converter"
	"user-service/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserInfoFromProto(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &user_v1.CreateResponse{
		Id: id,
	}, nil
}
