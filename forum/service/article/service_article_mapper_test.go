package article

import (
	"reflect"
	"schema/entity"
	"testing"
	"time"

	. "forum/model"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

func TestCommentResponseMapper(t *testing.T) {
	t.Run("when mapper works ok", func(t *testing.T) {
		comment := &entity.Comment{
			Body:      null.StringFrom("foo Body"),
			ArticleID: null.Uint64From(1),
			UserID:    null.Uint64From(1),
		}
		expected := &CommentResponse{
			Body: "foo Body",
		}
		actual := CommentResponseMapper(comment)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, want %v", actual, expected)
		}
	})
	t.Run("when created is valid", func(t *testing.T) {
		now := time.Now()
		comment := &entity.Comment{
			Body:      null.StringFrom("foo Body"),
			ArticleID: null.Uint64From(1),
			UserID:    null.Uint64From(1),
			CreatedAt: null.TimeFrom(now),
			UpdatedAt: null.TimeFrom(now),
		}
		expected := &CommentResponse{
			Body:      "foo Body",
			CreatedAt: now,
			UpdatedAt: now,
		}
		actual := CommentResponseMapper(comment)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, want %v", actual, expected)
		}
	})
}

func TestCommentListResponseMapper(t *testing.T) {
	t.Run("when mapper works ok", func(t *testing.T) {
		comments := []*entity.Comment{
			{
				ID:        1,
				Body:      null.StringFrom("foo Body"),
				ArticleID: null.Uint64From(1),
				UserID:    null.Uint64From(1),
			},
			{
				ID:        2,
				Body:      null.StringFrom("bar Body"),
				ArticleID: null.Uint64From(2),
				UserID:    null.Uint64From(2),
			},
		}

		actual := CommentListResponseMapper(comments)
		bodyList := []string{}
		for _, a := range actual.Comments {
			bodyList = append(bodyList, a.Body)
		}
		assert.Contains(t, bodyList, "foo Body", "foo Body should be in the list")
		assert.Contains(t, bodyList, "bar Body", "bar Body should be in the list")
	})
}
