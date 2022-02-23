package article

import (
	"forum/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestArticleCreateRequestBind(t *testing.T) {
	t.Run("When request is invalid, bind return error ", func(t *testing.T) {
		e := echo.New()
		e.Validator = utils.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/articles", strings.NewReader("invalid json"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &RequestArticle{}
		err := h.Bind(c, h)
		assert.IsType(t, &echo.HTTPError{}, err)
	})
}
