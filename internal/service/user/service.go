package user

import (
	"user-service/internal/client/db"
	"user-service/internal/repository"
	"user-service/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(userRepository repository.UserRepository, txManager db.TxManager) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
