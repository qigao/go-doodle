package model

import (
	"errors"
	"regexp"

	pb "gforum/grpc/gen/proto/forum/v1"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User is user model
type User struct {
	gorm.Model
	Username         string    `gorm:"unique_index;not null"`
	Email            string    `gorm:"unique_index;not null"`
	Password         string    `gorm:"not null"`
	Bio              string    `gorm:"not null"`
	Image            string    `gorm:"not null"`
	Follows          []User    `gorm:"many2many:follows;jointable_foreignkey:from_user_id;association_jointable_foreignkey:to_user_id"`
	FavoriteArticles []Article `gorm:"many2many:favorite_articles;"`
}

// Validate validates fields of user model
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Username,
			validation.Required,
			validation.Match(regexp.MustCompile("[a-zA-Z0-9]+")),
		),
		validation.Field(
			&u.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&u.Password,
			validation.Required,
		),
	)
}

// HashPassword makes password field crypted
func (u *User) HashPassword() error {
	if len(u.Password) == 0 {
		return errors.New("password should not be empty")
	}

	h, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(h)

	return nil
}

// CheckPassword checki user password correct
func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}

// ProtoUser generates proto user model from user
func (u *User) ProtoUser(token string) *pb.User {
	return &pb.User{
		Email:    u.Email,
		Token:    token,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
	}
}

// ProtoProfile generates proto profile model from user
func (u *User) ProtoProfile(following bool) *pb.Profile {
	return &pb.Profile{
		Username:  u.Username,
		Bio:       u.Bio,
		Image:     u.Image,
		Following: following,
	}
}
