package models

import (
	"errors"
	"strings"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"uniqueIndex;not null" json:"username" form:"username" valid:"required~Username is required"`
	Email        string        `gorm:"uniqueIndex;not null" json:"email" form:"email" valid:"required~Email is required,email~Invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~password minimum length at least a value of 6 characters"`
	Age          int           `gorm:"not null" json:"age" form:"age"`
	Photos       []Photo       `json:"photos" gorm:"constraint:OnDelete:CASCADE;"`
	SocialMedias []SocialMedia `json:"social_medias" gorm:"constraint:OnDelete:CASCADE;"`
	Comments     []Comment     `json:"comments" gorm:"constraint:OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	if !(u.Age > 8) {
		err = errors.New("minimum age is 9")
		return
	}

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(u.Username)) == 0 {
		err = errors.New("username is required")
		return
	}
	if len(strings.TrimSpace(u.Email)) == 0 {
		err = errors.New("email is required")
		return
	}

	if !(govalidator.IsEmail(u.Email)) {
		err = errors.New("invalid email format")
		return
	}

	err = nil
	return
}
