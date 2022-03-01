package article

import (
	"fmt"
	"forum/entity"
	"forum/mock/service"
	"forum/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newServiceArticleMock() *service.ServiceArticle {
	return &service.ServiceArticle{
		Mock: mock.Mock{},
	}
}

func TestArticleResource_GetArticle(t *testing.T) {
	t.Run("return Not-Found", func(t *testing.T) {
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(echo.GET, "/api/v1/", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug")
		c.SetParamNames("slug")
		c.SetParamValues("test-slug")
		serviceArticleMock := newServiceArticleMock()
		serviceArticleMock.On("FindArticle", mock.Anything).Return(nil, nil, nil, fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.GetArticle(c)
		require.NoError(t, err)
		assert.Error(t, echo.NewHTTPError(http.StatusOK, "error"))
	})
	t.Run("return OK", func(t *testing.T) {
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(echo.GET, "/api/v1/", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug")
		c.SetParamNames("slug")
		c.SetParamValues("test-slug")
		serviceArticleMock := newServiceArticleMock()
		serviceArticleMock.On("FindArticle", mock.Anything).Return(articleFoo, userFoo, []*entity.Tag{tagFoo}, nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.GetArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
