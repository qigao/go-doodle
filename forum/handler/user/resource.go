package user

import (
	"forum/handler"
	"forum/service"
	"forum/service/user"
	"github.com/labstack/echo/v4"
	"http/http_error"
	"net/http"

	"forum/entity"
)

// SignUp godoc
// @Summary Register a new user
// @Description Register a new user
// @ID sign-up
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body RegisterUser true "User info for registration"
// @Success 201 {object} userResponse
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /users [post]
func (h *Handler) SignUp(c echo.Context) error {
	var u entity.User
	req := &user.RegisterRequest{}
	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	if err := h.user.CreateUser(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	return c.JSON(http.StatusCreated, user.NewUserResponse(&u))
}

// Login godoc
// @Summary Login for existing user
// @Description Login for existing user
// @ID login
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body LoginRequest true "Credentials to use"
// @Success 200 {object} userResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /users/login [post]
func (h *Handler) Login(c echo.Context) error {
	req := &user.LoginRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	u, err := h.user.FindByEmail(req.User.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, http_error.AccessForbidden())
	}
	if !service.CheckPassword(u.Password, req.User.Password) {
		return c.JSON(http.StatusForbidden, http_error.AccessForbidden())
	}
	return c.JSON(http.StatusOK, user.NewUserResponse(u))
}

// CurrentUser godoc
// @Summary Get the current user
// @Description Gets the currently logged-in user
// @ID current-user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} Response
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /user [get]
func (h *Handler) CurrentUser(c echo.Context) error {
	u, err := h.user.FindByID(handler.UserIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, http_error.NotFound())
	}
	return c.JSON(http.StatusOK, user.NewUserResponse(u))
}

// UpdateUser godoc
// @Summary UpdateUser current user
// @Description UpdateUser user information for current user
// @ID update-user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body UpdateUser true "User details to update. At least **one** field is required."
// @Success 200 {object} userResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /user [put]
func (h *Handler) UpdateUser(c echo.Context) error {
	u, err := h.user.FindByID(handler.UserIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, http_error.NotFound())
	}
	req := user.NewUserUpdateRequest()
	req.Populate(u)
	if err := req.Bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	if err := h.user.UpdateUser(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, user.NewUserResponse(u))
}

// GetProfile godoc
// @Summary Get a profile
// @Description Get a profile of a user of the system. Auth is optional
// @ID get-profile
// @Tags profile
// @Accept  json
// @Produce  json
// @Param username path string true "Username of the profile to get"
// @Success 200 {object} Response
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /profiles/{username} [get]
func (h *Handler) GetProfile(c echo.Context) error {
	username := c.Param("username")
	u, err := h.user.FindByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, http_error.NotFound())
	}
	return c.JSON(http.StatusOK, user.NewProfileResponse(h.user, handler.UserIDFromToken(c), u))
}

// Follow godoc
// @Summary Follow a user
// @Description Follow a user by username
// @ID follow
// @Tags follow
// @Accept  json
// @Produce  json
// @Param username path string true "Username of the profile you want to follow"
// @Success 200 {object} profileResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /profiles/{username}/follow [post]
func (h *Handler) Follow(c echo.Context) error {
	followerID := handler.UserIDFromToken(c)
	username := c.Param("username")
	u, err := h.user.FindByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, http_error.NotFound())
	}
	if err := h.user.AddFollower(u, followerID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}

	return c.JSON(http.StatusOK, user.NewProfileResponse(h.user, handler.UserIDFromToken(c), u))
}

// Unfollow godoc
// @Summary Unfollow a user
// @Description Unfollow a user by username
// @ID unfollow
// @Tags follow
// @Accept  json
// @Produce  json
// @Param username path string true "Username of the profile you want to unfollow"
// @Success 201 {object} userResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /profiles/{username}/follow [delete]
func (h *Handler) Unfollow(c echo.Context) error {
	followerID := handler.UserIDFromToken(c)
	username := c.Param("username")
	u, err := h.user.FindByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, http_error.NotFound())
	}
	if err := h.user.RemoveFollower(u, followerID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, user.NewProfileResponse(h.user, handler.UserIDFromToken(c), u))
}
