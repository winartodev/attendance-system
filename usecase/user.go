package usecase

import (
	"context"

	"github.com/winartodev/attencande-system/ent"
	"github.com/winartodev/attencande-system/repository"
)

type UserUsecaseItf interface {
	Login(ctx context.Context, username string) (result *ent.User, err error)
	CreateUser(ctx context.Context, user ent.User) (err error)
}

type UserUsecase struct {
	UserRepository repository.UserRepositoryItf
}

func NewUserUsecase(userUsecase UserUsecase) UserUsecaseItf {
	return &UserUsecase{
		UserRepository: userUsecase.UserRepository,
	}
}
func (eu *UserUsecase) Login(ctx context.Context, username string) (result *ent.User, err error) {
	result, err = eu.UserRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return result, err
	}

	return result, err
}

func (eu *UserUsecase) CreateUser(ctx context.Context, user ent.User) (err error) {
	err = eu.UserRepository.CreateUserDB(ctx, user)
	if err != nil {
		return err
	}

	return err
}
