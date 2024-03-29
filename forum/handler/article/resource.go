package article

import (
	http_error "http/error"
	"net/http"
	"schema/entity"
	"strconv"

	"forum/handler"
	"forum/model"
	"forum/service/article"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
)

// GetArticle godoc
// @Summary Get an article
// @Description Get an article. Auth not required
// @ID get-article
// @Tags article
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the article to get"
// @Success 200 {object} singleArticleResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /articles/{slug} [get]
func (h *Handler) GetArticle(c echo.Context) error {
	slug := c.Param("slug")
	a, u, t, err := h.Service.FindArticle(slug)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get article")
		return c.JSON(http.StatusNotFound, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, article.SingleArticleResponseMapper(a, u, t))
}

// Articles godoc
// @Summary Get recent articles globally
// @Description Get most recent articles globally. Use query parameters to filter results. Auth is optional
// @ID get-articles
// @Tags article
// @Accept  json
// @Produce  json
// @Param tag query string false "Filter by tag"
// @Param author query string false "Filter by author (username)"
// @Param favorited query string false "Filter by favorites of a user (username)"
// @Param limit query integer false "Limit number of articles returned (default is 20)"
// @Param offset query integer false "Offset/skip number of articles (default is 0)"
// @Success 200 {object} articleListResponse
// @Failure 500 {object} utils.Error
// @Router /articles [get]
func (h *Handler) Articles(c echo.Context) error {
	tag := c.QueryParam("tag")
	author := c.QueryParam("author")

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		log.Error().Err(err).Msg("error parsing offset,set to 0")
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		log.Error().Err(err).Msg("error parsing limit,set to 20")
		limit = 20
	}

	articles, count, err := h.Service.FindArticles(tag, author, offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get articles")
		return c.JSON(http.StatusNotFound, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, article.SimpleArticleListMapper(articles, count))
}

// Feed godoc
// @Summary Get recent articles from users you follow
// @Description Get most recent articles from users you follow. Use query parameters to limit. Auth is required
// @ID feed
// @Tags article
// @Accept  json
// @Produce  json
// @Param limit query integer false "Limit number of articles returned (default is 20)"
// @Param offset query integer false "Offset/skip number of articles (default is 0)"
// @Success 200 {object} articleListResponse
// @Failure 401 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles/feed [get]
/* func (h *Handler) Feed(c echo.Context) error {
	var (
		articles []*entity.Article
		count    int64
	)

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	articles, count, err = h.article.ListFeed(handler.UserIDFromToken(c), offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, article.ArticleListResponseMapper(handler.UserIDFromToken(c), articles, count))
} */

// CreateArticle godoc
// @Summary Create an article
// @Description CreateUser an article. Auth is require
// @ID create-article
// @Tags article
// @Accept  json
// @Produce  json
// @Param article body CreateArticleRequest true "Article to create"
// @Success 201 {object} singleArticleResponse
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles [post]
func (h *Handler) CreateArticle(c echo.Context) error {
	var s model.SimpleArticle
	if err := c.Bind(&s); err != nil {
		log.Error().Err(err).Msg("error binding article")
		return c.JSON(http.StatusBadRequest, http_error.NewError(err))
	}
	a := populateSimpleArticle(&s)
	x := handler.UserIDFromToken(c)
	a.AuthorID = null.Uint64From(uint64(x))

	if err := h.Service.CreateArticle(a); err != nil {
		log.Error().Err(err).Msg("error inserting article")
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	return c.JSON(http.StatusCreated, handler.ResultOK())
}

// UpdateArticle godoc
// @Summary Update an article
// @Description UpdateUser an article. Auth is required
// @ID update-article
// @Tags article
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the article to update"
// @Param article body UpdateArticleRequest true "Article to update"
// @Success 200 {object} singleArticleResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles/{slug} [put]
func (h *Handler) UpdateArticle(c echo.Context) error {
	var s model.SimpleArticle
	slug := c.Param("slug")
	if err := c.Bind(&s); err != nil {
		log.Error().Err(err).Msg("error binding article")
		return c.JSON(http.StatusBadRequest, http_error.NotFound())
	}
	a := populateSimpleArticle(&s)
	x := handler.UserIDFromToken(c)
	a.AuthorID = null.Uint64From(uint64(x))
	if err := h.Service.UpdateArticle(slug, a); err != nil {
		log.Error().Err(err).Msg("error updating article")
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

// DeleteArticle godoc
// @Summary Delete an article
// @Description Delete an article. Auth is required
// @ID delete-article
// @Tags article
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the article to delete"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles/{slug} [delete]
func (h *Handler) DeleteArticle(c echo.Context) error {
	slug := c.Param("slug")
	err := h.Service.DeleteArticle(slug)
	if err != nil {
		log.Error().Err(err).Msg("error deleting article")
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

// AddComment godoc
// @Summary CreateUser a comment for an article
// @Description CreateUser a comment for an article. Auth is required
// @ID add-comment
// @Tags comment
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the article that you want to create a comment for"
// @Param comment body CreateCommentRequest true "Comment you want to create"
// @Success 201 {object} singleCommentResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles/{slug}/comments [post]
func (h *Handler) AddComment(c echo.Context) error {
	slug := c.Param("slug")
	var cm entity.Comment
	if err := c.Bind(&cm); err != nil {
		return c.JSON(http.StatusBadRequest, http_error.NewError(err))
	}
	if err := h.Service.AddCommentToArticle(slug, &cm); err != nil {
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{"result": "ok"})
}

// GetComments godoc
// @Summary Get the comments for an article
// @Description Get the comments for an article. Auth is optional
// @ID get-comments
// @Tags comment
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the article that you want to get comments for"
// @Success 200 {object} commentListResponse
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /articles/{slug}/comments [get]
func (h *Handler) GetComments(c echo.Context) error {
	slug := c.Param("slug")

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		log.Error().Err(err).Msg("error parsing offset,set to 0")
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		log.Error().Err(err).Msg("error parsing limit,set to 20")
		limit = 20
	}
	cms, err := h.Service.FindCommentsBySlug(slug, offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}

	return c.JSON(http.StatusOK, article.CommentListResponseMapper(cms))
}

// DeleteComment godoc
// @Summary Delete a comment for an article
// @Description Delete a comment for an article. Auth is required
// @ID delete-comments
// @Tags comment
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the article that you want to delete a comment for"
// @Param id path integer true "ID of the comment you want to delete"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles/{slug}/comments/{id} [delete]
func (h *Handler) DeleteComment(c echo.Context) error {
	x := c.Param("id")
	slug := c.Param("slug")
	id64, err := strconv.ParseUint(x, 10, 32)
	if err != nil {
		log.Error().Err(err).Msg("error parsing id")
		return c.JSON(http.StatusBadRequest, http_error.NewError(err))
	}

	if err := h.Service.DeleteCommentFromArticle(slug, id64); err != nil {
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

// Favorite godoc
// @Summary Favorite an article
// @Description Favorite an article. Auth is required
// @ID favorite
// @Tags favorite
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the article that you want to favorite"
// @Success 200 {object} singleArticleResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles/{slug}/favorite [post]
func (h *Handler) Favorite(c echo.Context) error {
	slug := c.Param("slug")
	x := handler.UserIDFromToken(c)
	err := h.Service.AddFavoriteArticleBySlug(slug, x)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

// Unfavorite godoc
// @Summary Unfavorite an article
// @Description Unfavorite an article. Auth is required
// @ID unfavorite
// @Tags favorite
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the article that you want to unfavorite"
// @Success 200 {object} singleArticleResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles/{slug}/favorite [delete]
func (h *Handler) Unfavorite(c echo.Context) error {
	slug := c.Param("slug")
	x := handler.UserIDFromToken(c)
	err := h.Service.RemoveFavoriteArticleBySlug(slug, x)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error removing favorite")
		return c.JSON(http.StatusInternalServerError, http_error.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

// Tags godoc
// @Summary Get tags
// @Description Get tags
// @ID tags
// @Tags tag
// @Accept  json
// @Produce  json
// @Success 201 {object} tagListResponse
// @Failure 400 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /tags [get]
func (h *Handler) Tags(c echo.Context) error {
	tags, err := h.Service.GetAllTags()
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, article.TagListResponseMapper(tags))
}

// AddTagToArticle godoc
// @Summary tag an article
// @Description tag an article. Auth is required
// @ID tags
// @Tags tag
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the article to tag"
// @Param tag  path string true "existing tag to apply"
// @Param article body UpdateArticleRequest true "Article to update"
// @Success 200 {object} singleArticleResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles/{slug}/{tag} [put]
func (h *Handler) AddTagToArticle(c echo.Context) error {
	slug := c.Param("slug")
	tag := c.Param("tag")
	if err := h.Service.AddTagToArticle(slug, []string{tag}); err != nil {
		return c.JSON(http.StatusBadRequest, http_error.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}
