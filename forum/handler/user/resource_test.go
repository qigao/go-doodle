package user

import (
	"forum/service/user"
	"forum/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
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
		h := &user.RequestUser{}
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
		h := &user.RequestUser{}
		// Assertions
		err := h.Bind(c)
		assert.Error(t, err)
		assert.Equal(t, "Key: 'RequestLogin.User.Email' Error:Field validation for 'Email' failed on the 'required' tag", err.Error())
	})
	t.Run("When request json is invalid,Bind return Error ", func(t *testing.T) {
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/login", strings.NewReader("invalid json"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &user.RequestUser{}
		// Assertions
		err := h.Bind(c)
		assert.IsType(t, &echo.HTTPError{}, err)
	})
}
