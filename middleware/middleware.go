package middleware

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/winartodev/attencande-system/helper"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("Token")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed{
				Status:  http.StatusText(http.StatusUnauthorized),
				Message: helper.ErrNoToken.Error(),
			})
		}

		token := cookie.Value
		claims, err := helper.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed{
				Status:  http.StatusText(http.StatusUnauthorized),
				Message: helper.ErrNoToken.Error(),
			})
		}

		c.Set("role", claims.Role)

		return next(c)
	}
}
