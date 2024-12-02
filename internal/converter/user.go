package converter

import (
	"user-service/internal/model"
	"user-service/pkg/user_v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoUser(user *model.User) *user_v1.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &user_v1.User{
		Id:        user.ID,
		Info:      ToProtoUserInfo(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToProtoUserInfo(info *model.UserInfo) *user_v1.UserInfo {
	return &user_v1.UserInfo{
		FirstName:   info.FirstName,
		LastName:    info.LastName,
		Password:    info.Password,
		PhoneNumber: info.PhoneNumber,
		Email:       info.Email,
	}
}

func ToUserInfoFromProto(info *user_v1.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		FirstName:   info.FirstName,
		LastName:    info.LastName,
		Password:    info.Password,
		PhoneNumber: info.PhoneNumber,
		Email:       info.Email,
	}
}
