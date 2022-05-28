package article

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"schema/entity"
	"strings"
	"testing"

	"forum/mock/service"
	"forum/model"
	"forum/utils"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestArticleResource_GetArticle(t *testing.T) {
	t.Run("When return Not-Found", func(t *testing.T) {
		rec, c := echoFindArticleSetup()
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("FindArticle", mock.Anything).Return(nil, nil, nil, fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.GetArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
	t.Run("When return OK", func(t *testing.T) {
		rec, c := echoFindArticleSetup()
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("FindArticle", mock.Anything).Return(articleFoo, userFoo, []*entity.Tag{tagFoo}, nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.GetArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func echoFindArticleSetup() (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	e.Validator = utils.NewValidator()
	req := httptest.NewRequest(echo.GET, "/api/v1/articles", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:slug")
	c.SetParamNames("slug")
	c.SetParamValues("test-slug")
	return rec, c
}

func TestArticleResource_GetArticles(t *testing.T) {
	t.Run("When return OK", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = utils.NewValidator()
		q := make(url.Values)
		q.Set("tag", "newTag")
		q.Set("author", "newAuthor")
		q.Set("offset", "0")
		q.Set("limit", "10")
		req := httptest.NewRequest(echo.GET, "/api/v1/articles?"+q.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("FindArticles", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Article{articleFoo}, int64(1), nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.Articles(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("When offset/limit is nil, return OK", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = utils.NewValidator()
		q := make(url.Values)
		q.Set("tag", "newTag")
		q.Set("author", "newAuthor")
		q.Set("offset", "0")
		q.Set("limit", "10")
		req := httptest.NewRequest(echo.GET, "/api/v1/articles?"+q.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("FindArticles", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Article{articleFoo}, int64(1), nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.Articles(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("when find articles return error", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = utils.NewValidator()
		q := make(url.Values)
		q.Set("tag", "newTag")
		q.Set("author", "newAuthor")
		q.Set("offset", "0")
		q.Set("limit", "10")
		req := httptest.NewRequest(echo.GET, "/api/v1/articles?"+q.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("FindArticles", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.Articles(c)
		require.NoError(t, err)
		assert.Error(t, echo.NewHTTPError(http.StatusNotFound, "error"))
	})
}

func TestArticleResource_CreateArticle(t *testing.T) {
	t.Run("when create article return ok", func(t *testing.T) {
		// When
		simple := &model.SimpleArticle{
			Title:       "test",
			Slug:        "test",
			Description: "test",
			Body:        "test",
		}
		data, err := json.Marshal(simple)
		reader := bytes.NewReader(data)
		require.NoError(t, err)
		// Setup
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/", reader)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("CreateArticle", mock.Anything).Return(nil)
		handler := NewArticleHandler(serviceArticleMock)
		err = handler.CreateArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
	})
	t.Run("when bind error", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(echo.POST, "/api/v1/", strings.NewReader("invalid json"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.CreateArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("when create article return error", func(t *testing.T) {
		// When
		simple := &model.SimpleArticle{
			Title:       "test",
			Slug:        "test",
			Description: "test",
			Body:        "test",
		}
		data, err := json.Marshal(simple)
		reader := bytes.NewReader(data)
		require.NoError(t, err)
		// Setup
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/", reader)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("CreateArticle", mock.Anything).Return(fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err = handler.CreateArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestArticleResource_UpdateArticle(t *testing.T) {
	t.Run("when update return OK", func(t *testing.T) {
		// When
		simple := &model.SimpleArticle{
			Title:       "test",
			Slug:        "test",
			Description: "test",
			Body:        "test",
		}
		data, err := json.Marshal(simple)
		reader := bytes.NewReader(data)
		require.NoError(t, err)
		// Setup
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(echo.PUT, "/api/v1/", reader)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("UpdateArticle", mock.Anything, mock.Anything).Return(nil)
		handler := NewArticleHandler(serviceArticleMock)
		err = handler.UpdateArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("when bind error", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(echo.POST, "/api/v1/", strings.NewReader("invalid json"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.UpdateArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("when update return error", func(t *testing.T) {
		// When
		simple := &model.SimpleArticle{
			Title:       "test",
			Slug:        "test",
			Description: "test",
			Body:        "test",
		}
		data, err := json.Marshal(simple)
		reader := bytes.NewReader(data)
		require.NoError(t, err)
		// Setup
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(echo.PUT, "/api/v1/", reader)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("UpdateArticle", mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err = handler.UpdateArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestArticleResource_DeleteArticle(t *testing.T) {
	t.Run("when delete return OK", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("DeleteArticle", mock.Anything).Return(nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.DeleteArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("when delete return error", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("DeleteArticle", mock.Anything).Return(fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.DeleteArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestArticleResource_AddComment(t *testing.T) {
	t.Run("when add comment return OK", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("AddCommentToArticle", mock.Anything, mock.Anything).Return(nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.AddComment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
	})
	t.Run("when add comment return error", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("AddCommentToArticle", mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.AddComment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
	t.Run("when bind error", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(echo.POST, "/api/v1/", strings.NewReader("invalid json"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.AddComment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestArticleResource_GetComments(t *testing.T) {
	t.Run("when get comments return OK", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("FindCommentsBySlug", mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Comment{commentFoo}, nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.GetComments(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("when FindCommentsBySlug return error", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("FindCommentsBySlug", mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.GetComments(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestArticleResource_DeleteComent(t *testing.T) {
	t.Run("when delete comment from article return ok", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug/comments/:id")
		c.SetParamNames("slug", "id")
		c.SetParamValues("test-slug", "1")
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("DeleteCommentFromArticle", mock.Anything, mock.Anything).Return(nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.DeleteComment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("when comments id is invalid", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug/comments/:id")
		c.SetParamNames("id", "slug")
		c.SetParamValues("invalid", "test-slug")
		serviceArticleMock := service.NewIServiceArticle(t)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.DeleteComment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("when delete comment from article return error", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug/comments/:id")
		c.SetParamNames("slug", "id")
		c.SetParamValues("test-slug", "1")
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("DeleteCommentFromArticle", mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.DeleteComment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestArticleResource_Favorite(t *testing.T) {
	t.Run("when favorite article return ok", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug/favorite")
		c.SetParamNames("slug")
		c.SetParamValues("test-slug")
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("AddFavoriteArticleBySlug", mock.Anything, mock.Anything).Return(nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.Favorite(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("when favorite article return error", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug/favorite")
		c.SetParamNames("slug")
		c.SetParamValues("test-slug")
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("AddFavoriteArticleBySlug", mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.Favorite(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestArticleResource_UnFavorite(t *testing.T) {
	t.Run("when unFavorite article return ok", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug/favorite")
		c.SetParamNames("slug")
		c.SetParamValues("test-slug")
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("RemoveFavoriteArticleBySlug", mock.Anything, mock.Anything).Return(nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.Unfavorite(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("when unFavorite article return error", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug/favorite")
		c.SetParamNames("slug")
		c.SetParamValues("test-slug")
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("RemoveFavoriteArticleBySlug", mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.Unfavorite(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestArticleResource_Tags(t *testing.T) {
	t.Run("when get all tags return ok", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("GetAllTags").Return([]*entity.Tag{tagBar, tagFoo}, nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.Tags(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("when get all tags return error", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("GetAllTags").Return(nil, fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.Tags(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestArticleResource_AddTagToArticle(t *testing.T) {
	t.Run("when add tag to article return ok", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug/:tag")
		c.SetParamNames("slug", "tag")
		c.SetParamValues("test-slug", "new-tag")
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("AddTagToArticle", mock.Anything, mock.Anything).Return(nil)
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.AddTagToArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("when add tag to article return error", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles/:slug/:tag")
		c.SetParamNames("slug", "tag")
		c.SetParamValues("test-slug", "new-tag")
		serviceArticleMock := service.NewIServiceArticle(t)
		serviceArticleMock.On("AddTagToArticle", mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
		handler := NewArticleHandler(serviceArticleMock)
		err := handler.AddTagToArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
