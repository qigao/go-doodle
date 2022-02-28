package article

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"forum/model"
	"forum/utils"
)

func TestArticleResource_CreateArticle(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(`{"title": "hello", "content": "world"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+utils.GenerateJWT(utils.JWTConfig, "user", "1"))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		articleRepo := &mockArticleRepo{
			createFunc: func(article *model.Article) error {
				article.ID = 1
				return nil
			},
		}
		articleService := &mockArticleService{
			createFunc: func(article *model.Article) (*model.Article, error) {
				article.ID = 1
				return article, nil
			},
		}
		h := &Handler{
			articleRepo:    articleRepo,
			articleService: articleService,
		}

		err := h.CreateArticle(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `{"id":1,"title":"hello","content":"world","author_id":1}`, rec.Body.String())
	})
}
