package article

import (
	"forum/entity"
	"forum/handler"
	"forum/model"
	"github.com/labstack/echo/v4"
)

type singleCommentResponse struct {
	Comment *model.CommentResponse `json:"comment"`
}

type commentListResponse struct {
	Comments []model.CommentResponse `json:"comments"`
}

func NewCommentResponse(c echo.Context, cm *entity.Comment) *singleCommentResponse {
	comment := new(model.CommentResponse)
	comment.ID = cm.ID
	comment.Body = cm.Body
	comment.CreatedAt = cm.CreatedAt
	comment.UpdatedAt = cm.UpdatedAt
	comment.Author.Username = cm.User.Username
	comment.Author.Image = cm.User.Image
	comment.Author.Bio = cm.User.Bio
	comment.Author.Following = cm.User.FollowedBy(handler.UserIDFromToken(c))
	return &singleCommentResponse{comment}
}

func NewCommentListResponse(c echo.Context, comments []entity.Comment) *commentListResponse {
	r := new(commentListResponse)
	cr := model.CommentResponse{}
	r.Comments = make([]model.CommentResponse, 0)
	for _, i := range comments {
		cr.ID = i.ID
		cr.Body = i.Body
		cr.CreatedAt = i.CreatedAt
		cr.UpdatedAt = i.UpdatedAt
		cr.Author.Username = i.User.Username
		cr.Author.Image = i.User.Image
		cr.Author.Bio = i.User.Bio
		cr.Author.Following = i.User.FollowedBy(handler.UserIDFromToken(c))

		r.Comments = append(r.Comments, cr)
	}
	return r
}
