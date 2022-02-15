package article

import (
	"forum/entity"
	"forum/handler"
	"github.com/labstack/echo/v4"
)

type CommentRequest struct {
	Comment struct {
		Body string `json:"body" validate:"required"`
	} `json:"comment"`
}

func (r *CommentRequest) Bind(c echo.Context, cm *entity.Comment) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	cm.Body = r.Comment.Body
	cm.UserID = handler.UserIDFromToken(c)
	return nil
}
