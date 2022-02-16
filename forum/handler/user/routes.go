package user

import (
	"forum/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v *echo.Group) {
	jwtMiddleware := utils.JWT(utils.JWTSecret)
	guestUsers := v.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := v.Group("/user", jwtMiddleware)
	user.GET("", h.CurrentUser)
	user.PUT("", h.UpdateUser)

	profiles := v.Group("/profiles", jwtMiddleware)
	profiles.GET("/:username", h.GetProfile)
	profiles.POST("/:username/follow", h.Follow)
	profiles.DELETE("/:username/follow", h.Unfollow)
}
