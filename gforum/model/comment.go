package model

import (
	"fmt"

	pb "gforum/grpc/gen/proto/forum/v1"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

// Comment model
type Comment struct {
	gorm.Model
	Body      string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	Author    User   `gorm:"foreignkey:UserID"`
	ArticleID uint   `gorm:"not null"`
	Article   Article
}

// Validate validates fields of comment model
func (c Comment) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(
			&c.Body,
			validation.Required,
		),
	)
}

// ProtoComment generates proto comment model from article
func (c *Comment) ProtoComment() *pb.Comment {
	return &pb.Comment{
		Id:        fmt.Sprintf("%d", c.ID),
		Body:      c.Body,
		CreatedAt: c.CreatedAt.Format(ISO8601),
		UpdatedAt: c.UpdatedAt.Format(ISO8601),
	}
}
