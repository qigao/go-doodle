package article

import (
	"forum/entity"
	"forum/model"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
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
