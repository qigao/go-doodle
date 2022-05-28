package user

import (
	"http/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v *echo.Group) {
	jwtMiddleware := utils.JWT(utils.JWTSecret)
	guestUsers := v.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := v.Group("/user", jwtMiddleware)
	user.PUT("", h.UpdateUser)
}
