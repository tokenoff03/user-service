package service

import (
	"context"
	"user-service/internal/model"
)

type UserService interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
}
