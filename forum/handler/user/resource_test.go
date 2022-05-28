package user

import (
	"fmt"
	"forum/mock/service"
	"forum/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin_Bind(t *testing.T) {
	jsonUser := `{"user":{"email":"alice@realworld.io","password":"secret"}}`
	const ApiLogin = "/api/v1/login"
	t.Run("When Bind return OK ", func(t *testing.T) {
		rec, c := echoSetup(http.MethodPost, ApiLogin, jsonUser)
		serviceUserMock := service.NewIServiceUser(t)
		serviceUserMock.On("CheckUser", mock.Anything).Return(nil)
		handler := NewUserHandler(serviceUserMock)
		err := handler.Login(c)
		require.NoError(t, err)
		// Assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("When CheckUser return Error", func(t *testing.T) {
		rec, c := echoSetup(http.MethodPost, ApiLogin, jsonUser)
		serviceUserMock := service.NewIServiceUser(t)
		serviceUserMock.On("CheckUser", mock.Anything).Return(fmt.Errorf("error"))
		handler := NewUserHandler(serviceUserMock)
		err := handler.Login(c)
		// Assertions
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
	t.Run("When request json is invalid,Bind return Error ", func(t *testing.T) {
		rec, c := echoSetup(http.MethodPost, ApiLogin, "invalid json")
		serviceUserMock := service.NewIServiceUser(t)
		handler := NewUserHandler(serviceUserMock)
		// Assertions
		err := handler.Login(c)
		// Assertions
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func echoSetup(method string, url string, jsonUser string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	e.Validator = utils.NewValidator()
	req := httptest.NewRequest(method, url, strings.NewReader(jsonUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}
