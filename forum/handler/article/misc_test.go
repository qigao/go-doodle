package article

import (
	"forum/entity"
	"forum/model"
	"reflect"
	"testing"

	"github.com/volatiletech/null/v8"
)

func Test_populateSingleArticle(t *testing.T) {
	type args struct {
		s *model.SimpleArticle
	}
	tests := []struct {
		name string
		args args
		want *entity.Article
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := populateSingleArticle(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("populateSingleArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	userFoo = &entity.User{
		Username: "foo",
		Email:    "foo@foo.com",
		Password: "123456",
	}
	userBar = &entity.User{
		Username: "bar",
		Email:    "bar@bar.com",
		Password: "123456",
	}
	articleFoo = &entity.Article{
		Title:       "foo Title",
		Description: null.NewString("foo Description", false),
		Body:        null.NewString("foo Body", false),
		Slug:        "foo-slug",
	}
	articleBar = &entity.Article{
		Title:       "foo Title",
		Description: null.NewString("foo Description", false),
		Body:        null.NewString("foo Body", false),
		Slug:        "foo-slug",
	}
	commentFoo = &entity.Comment{
		Body:      null.StringFrom("foo Body"),
		ArticleID: null.Uint64From(1),
		UserID:    null.Uint64From(1),
	}
	tagFoo = &entity.Tag{
		Tag: null.StringFrom("foo"),
	}
	tagBar = &entity.Tag{
		Tag: null.StringFrom("bar"),
	}
)
