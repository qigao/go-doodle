package user

import (
	"forum/handler"
	"forum/model"
	"forum/service/user"
	"http/http_error"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
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
	var reg model.RegisterUser
	if err := c.Bind(&reg); err != nil {
		log.Error().Err(err).Msg("Error binding request")
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	if err := h.Service.CreateUser(&reg); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	return c.JSON(http.StatusCreated, handler.ResultOK())
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
	var req model.LoginUser
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	err := h.Service.CheckUser(&req)
	if err != nil {
		return c.JSON(http.StatusNotFound, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, handler.ResultOK())
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
	uId := handler.UserIDFromToken(c)
	req := &user.UserProfile{Repo: h.userRepo}
	userInfo, err := req.GetUserByID(uId)
	if err != nil {
		return c.JSON(http.StatusNotFound, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, user.NewUserResponse(userInfo))
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
	uid := handler.UserIDFromToken(c)
	req := &user.UpdateRequest{Repo: h.userRepo}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	if err := req.FindThenUpdateUser(uid); err != nil {
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
	req := &user.UserProfile{Repo: h.userRepo}
	u, err := req.GetUserByUsername(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, user.NewProfileResponse(h.userRepo, handler.UserIDFromToken(c), u))
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
	req := &user.FollowRequest{Repo: h.userRepo}
	if err := req.FllowUserByUserName(followerID, username); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}

	return c.JSON(http.StatusOK, user.NewProfileResponse(h.userRepo, handler.UserIDFromToken(c), u))
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
	uid := handler.UserIDFromToken(c)
	userName := c.Param("username")
	req := &user.FollowRequest{Repo: h.userRepo}
	if err := req.UnFllowUserByUserName(uid, userName); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, user.NewProfileResponse(h.userRepo, handler.UserIDFromToken(c), u))
}
