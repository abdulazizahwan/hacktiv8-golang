package models

import (
	"errors"
	"strings"


	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint   `json:"UserID"`
	User    User   `json:"user"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Photo   Photo  `json:"photo"`
	Message string `gorm:"not null" json:"message" form:"message"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB)(err error){
	if len(strings.TrimSpace(c.Message))==0{
		err=errors.New("message is required")
		return
	}
	err=nil
	return
}