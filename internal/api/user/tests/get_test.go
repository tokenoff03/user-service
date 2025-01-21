package tests

import (
	"context"
	"database/sql"
	"testing"
	"time"
	"user-service/internal/api/user"
	"user-service/internal/model"
	"user-service/internal/service"
	serviceMocks "user-service/internal/service/mocks"
	"user-service/pkg/user_v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *user_v1.GetRequest
	}

	var (
		id  = gofakeit.Int64()
		ctx = context.Background()
		req = &user_v1.GetRequest{
			Id: id,
		}
		mc          = minimock.NewController(t)
		firstName   = gofakeit.Name()
		lastName    = gofakeit.LastName()
		password    = gofakeit.Password(true, true, true, false, false, 8)
		phoneNumber = gofakeit.Phone()
		email       = gofakeit.Email()
		time        = time.Now()
		createdAt   = timestamppb.New(time)
		updated_at  = timestamppb.New(time)
		res         = &user_v1.GetResponse{
			User: &user_v1.User{
				Id: id,
				Info: &user_v1.UserInfo{
					FirstName:   firstName,
					LastName:    lastName,
					Password:    password,
					PhoneNumber: phoneNumber,
					Email:       email,
				},
				CreatedAt: createdAt,
				UpdatedAt: updated_at,
			},
		}
		userModel = &model.User{
			ID: id,
			Info: &model.UserInfo{
				FirstName:   firstName,
				LastName:    lastName,
				Password:    password,
				PhoneNumber: phoneNumber,
				Email:       email,
			},
			CreatedAt: time,
			UpdatedAt: sql.NullTime{time, true},
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *user_v1.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success test",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(userModel, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMockFunc := tt.userServiceMock(mc)
			api := user.NewImplementation(userServiceMockFunc)
			res, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
