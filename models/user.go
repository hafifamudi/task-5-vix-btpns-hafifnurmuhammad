package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"not null" json:"username" form:"username" valid:"required~Username is required"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid Email"`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password minimum is 6 characters"`
	Photo     []Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
	CreatedAt *time.Time `json:",omitempty"`
	UpdatedAt *time.Time `json:",omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errVal := govalidator.ValidateStruct(u)

	if errVal != nil {
		return errVal
	}

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errVal := govalidator.ValidateStruct(u)

	if errVal != nil {
		return errVal
	}

	return errVal
}
