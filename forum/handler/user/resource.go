package user

import (
	"forum/handler"
	"forum/model"
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
		return c.JSON(http.StatusBadRequest, http_error.NewError(err))
	}
	err := h.Service.CheckUser(&req)
	if err != nil {
		return c.JSON(http.StatusNotFound, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, handler.ResultOK())
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
	var s model.ProfileType
	if err := c.Bind(&s); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	u, err := h.Service.GetUserByID(uid)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	if err = h.Service.UpdateUser(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, handler.ResultOK())
}
