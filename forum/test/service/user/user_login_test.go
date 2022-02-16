package user

import (
	user2 "forum/service/user"
	"forum/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserLogin_Bind(t *testing.T) {
	t.Run("When Bind return OK ", func(t *testing.T) {
		jsonUser := `{"user":{"email":"alice@realworld.io","password":"secret"}}`
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/login", strings.NewReader(string(jsonUser)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &user2.LoginRequest{}
		// Assertions
		assert.NoError(t, h.Bind(c))
	})
	t.Run("When validate return Error ", func(t *testing.T) {
		jsonUser := `{"user":{"emaill":"alice@realworld.io","password":"secret"}}`
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/login", strings.NewReader(string(jsonUser)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &user2.LoginRequest{}
		// Assertions
		err := h.Bind(c)
		assert.Error(t, err)
		assert.Equal(t, "Key: 'LoginRequest.User.Email' Error:Field validation for 'Email' failed on the 'required' tag", err.Error())
	})
	t.Run("When Bind return Error ", func(t *testing.T) {
		jsonUser := `{"user":{"email":"alice@realworld.io","password":"secret\"}}`
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/login", strings.NewReader(string(jsonUser)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &user2.LoginRequest{}
		// Assertions
		err := h.Bind(c)
		assert.Error(t, err)
	})
}
