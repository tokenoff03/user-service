package tests

import (
	"context"
	"fmt"
	"log"
	"testing"
	"user-service/internal/api/user"
	"user-service/internal/model"
	"user-service/internal/service"
	serviceMocks "user-service/internal/service/mocks"
	"user-service/pkg/user_v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"
	"github.com/tokenoff03/authentication-service/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCreate(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *user_v1.CreateRequest
	}

	var (
		ctx         = context.Background()
		mc          = minimock.NewController(t)
		id          = gofakeit.Int64()
		firstName   = gofakeit.Name()
		lastName    = gofakeit.LastName()
		password    = gofakeit.Password(true, true, true, false, false, 8)
		phoneNumber = gofakeit.Phone()
		email       = gofakeit.Email()
		req         = &user_v1.CreateRequest{
			Info: &user_v1.UserInfo{
				FirstName:   firstName,
				LastName:    lastName,
				Password:    password,
				PhoneNumber: phoneNumber,
				Email:       email,
			},
		}

		res = &user_v1.CreateResponse{
			Id: id,
		}

		serviceErr = fmt.Errorf("service error")

		info = &model.UserInfo{
			FirstName:   firstName,
			LastName:    lastName,
			Password:    password,
			PhoneNumber: phoneNumber,
			Email:       email,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *user_v1.CreateResponse
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
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			conn, err := grpc.NewClient(
				fmt.Sprintf(":%d", "50059"),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
			)
			if err != nil {
				log.Fatalf("failed to dial GRPC client: %v", err)
			}
			userServiceMockFunc := tt.userServiceMock(mc)
			api := user.NewImplementation(userServiceMockFunc, auth_v1.NewAuthV1Client(conn))
			res, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}

}
