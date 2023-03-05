package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Title     string `json:"title" form:"title" valid:"required~Photo Title is Required"`
	Caption   string `json:"caption" form:"caption"`
	PhotoURL  string `json:"photo_url" form:"photo_url" valid:"required~Photo URL is Required"`
	UserID    int    `json:"user_id"`
	User      *User
	CreatedAt *time.Time `json:",omitempty"`
	UpdatedAt *time.Time `json:",omitempty"`
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errVal := govalidator.ValidateStruct(u)

	if errVal != nil {
		return errVal
	}

	return nil
}

func (u *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errVal := govalidator.ValidateStruct(u)

	if errVal != nil {
		return errVal
	}

	return errVal
}
