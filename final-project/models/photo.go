package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~Photo title is required"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoURL string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo url is required"`
	UserID   uint      `json:"user_id" form:"user_id"`
	User     User      `json:"user"`
	Comments []Comment `json:"Comments" gorm:"constraint:OnDelete:CASCADE;"`
}

type PhotoValidation struct {
	Title    string `json:"title" form:"title" valid:"required~Photo title is required"`
	PhotoURL string `json:"photo_url" form:"photo_url" valid:"required~Photo url is required"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(p.Title)) == 0 {
		err = errors.New("photo title is required")
		return
	}
	if len(strings.TrimSpace(p.PhotoURL)) == 0 {
		err = errors.New("photo url is required")
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(p.Title)) == 0 {
		err = errors.New("photo title is required")
		return
	}
	if len(strings.TrimSpace(p.PhotoURL)) == 0 {
		err = errors.New("photo url is required")
		return
	}

	err = nil
	return
}
