package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/winartodev/attencande-system/helper"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		fmt.Println(authorizationHeader)
		if !strings.Contains(authorizationHeader, "Bearer") {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed{
				Status:  http.StatusText(http.StatusUnauthorized),
				Message: helper.ErrNoToken.Error(),
			})
		}
		token := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		_, err := helper.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed{
				Status:  http.StatusText(http.StatusUnauthorized),
				Message: helper.ErrNoToken.Error(),
			})
		}

		return next(c)
	}
}
