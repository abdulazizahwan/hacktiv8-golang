package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaURL string `gorm:"not null" form:"social_media_url" json:"social_media_url"`
	UserID         uint   `json:"user_id"`
	User           User   `json:"user"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(s.Name)) == 0 {
		err = errors.New("social media name is required")
		return
	}
	if len(strings.TrimSpace(s.SocialMediaURL)) == 0 {
		err = errors.New("social media url is required")
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(s.Name)) == 0 {
		err = errors.New("social media name is required")
		return
	}
	if len(strings.TrimSpace(s.SocialMediaURL)) == 0 {
		err = errors.New("social media url is required")
		return
	}

	err = nil
	return
}
