package repository

import (
	"context"

	"github.com/winartodev/attencande-system/ent"
	"github.com/winartodev/attencande-system/ent/user"
)

type UserRepositoryItf interface {
	CreateUserDB(ctx context.Context, user ent.User) (err error)
	GetUserByUsername(ctx context.Context, username string) (result *ent.User, err error)
}

type UserRepository struct {
	Client *ent.Client
}

func NewUserRepository(userRepository UserRepository) UserRepositoryItf {
	return &UserRepository{
		Client: userRepository.Client,
	}
}

func (er *UserRepository) CreateUserDB(ctx context.Context, user ent.User) (err error) {
	_, err = er.Client.User.
		Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetRole(user.Role).
		Save(ctx)

	if err != nil {
		return err
	}

	return err
}

func (er *UserRepository) GetUserByUsername(ctx context.Context, username string) (result *ent.User, err error) {
	result, err = er.Client.User.
		Query().
		Select(user.FieldID, user.FieldUsername, user.FieldPassword, user.FieldRole).
		Where(user.Username(username)).
		Only(ctx)
	if err != nil {
		return result, err
	}

	return result, err
}
