package user

import (
	"context"
	"user-service/internal/converter"
	"user-service/pkg/user_v1"

	"github.com/tokenoff03/lib_ad1lek/pkg/sys"
	"github.com/tokenoff03/lib_ad1lek/pkg/sys/codes"
	"github.com/tokenoff03/lib_ad1lek/pkg/sys/validate"
)

func (i *Implementation) Get(ctx context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	err := validate.Validate(
		ctx,
		validateID(req.GetId()),
		otherValidateID(req.GetId()),
	)

	if err != nil {
		return nil, err
	}

	if req.GetId() > 100 {
		return nil, sys.NewCommonError("id must be less than 100", codes.ResourceExhausted)
	}

	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &user_v1.GetResponse{
		User: converter.ToProtoUser(user),
	}, nil
}
