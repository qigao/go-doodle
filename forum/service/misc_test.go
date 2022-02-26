package service

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestCheckPassword(t *testing.T) {
	type args struct {
		plain      string
		hashedPass string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "when plain is empty", args: args{plain: "", hashedPass: ""}, wantErr: true},
		{name: "when plain is not empty", args: args{plain: "123456", hashedPass: "$2a$10$B65SchLWy/AqA75Oap8jO.ZJGTtF40/6elzX1mYv0W/0K.yQQw7WW"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckPassword(tt.args.plain, tt.args.hashedPass); (err != nil) != tt.wantErr {
				t.Errorf("CheckPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashPassword(t *testing.T) {
	type args struct {
		plain string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "when plain is empty", args: args{plain: ""}, want: "", wantErr: true},
		{name: "when plain is not empty", args: args{plain: "123456"}, want: "$2a$10$B65SchLWy/AqA75Oap8jO.ZJGTtF40/6elzX1mYv0W/0K.yQQw7WW", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.plain)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err == nil {
				msg := fmt.Sprintf("HashPassword() error = %v, wantErr %v got %s", err, tt.wantErr, got)
				assert.NilError(t, err, msg)
			}
		})
	}
}
