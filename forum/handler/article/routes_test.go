package article

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHandler_Register(t *testing.T) {
	type args struct {
		v *echo.Group
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Register(tt.args.v)
		})
	}
}
