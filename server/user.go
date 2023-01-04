package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/winartodev/attencande-system/ent"
	"github.com/winartodev/attencande-system/helper"
	"github.com/winartodev/attencande-system/usecase"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecaseItf
}

func NewUserUsecase(userHandler UserHandler) UserHandler {
	return UserHandler{
		UserUsecase: userHandler.UserUsecase,
	}
}

// RegisterHandler is used to register new user
func (uh *UserHandler) RegisterHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var user ent.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusUnprocessableEntity),
			Message: err.Error(),
		})
	}

	password := helper.HashPassword(user.Password)
	user.Password = password

	err = uh.UserUsecase.CreateUser(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "User success created",
	})
}

// LoginHandler will get data by username and password
func (uh *UserHandler) LoginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	ctx := c.Request().Context()

	result, err := uh.UserUsecase.Login(ctx, username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	ok, err := helper.VerifyPassword(password, result.Password)
	if !ok {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusUnauthorized),
			Message: err.Error(),
		})
	}

	token, err := helper.GenerateToken(username, result.Role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success Loggedin",
		Data:    token,
	})
}

func (uh *UserHandler) LogoutHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success logout",
	})
}
