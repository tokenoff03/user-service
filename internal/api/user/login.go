package user

import (
	"context"
	"user-service/pkg/user_v1"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/tokenoff03/authentication-service/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *user_v1.LoginRequest) (*user_v1.LoginResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "login")
	defer span.Finish()

	span.SetTag("email", req.GetEmail())
	refreshToken, err := i.authServiceClient.Login(ctx, &auth_v1.LoginRequest{Email: req.GetEmail(), Password: req.GetPassword()})

	if err != nil {
		return nil, errors.WithMessage(err, "loginning")
	}

	return &user_v1.LoginResponse{
		RefreshToken: refreshToken.RefreshToken,
	}, nil
}
