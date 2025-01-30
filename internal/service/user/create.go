package user

import (
	"context"
	"user-service/internal/model"
	"user-service/internal/utils"
)

func (s *serv) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	var id int64
	hashedPassword, err := utils.HashPassword(info.Password)
	if err != nil {
		return 0, err
	}
	info.Password = hashedPassword

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userRepository.Get(ctx, id) //Можно было сделать Лог для вывода создания пользователя!
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}
