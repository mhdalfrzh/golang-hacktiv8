package models

import (
	"final-project/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Username     string     `gorm:"not null; uniqueIndex" json:"username"`
	Email        string     `gorm:"not null; uniqueIndex" json:"email" valid:"required~Your emails is required, email~Invalid email format"`
	Password     string     `gorm:"not null" json:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          uint       `gorm:"not null" json:"age" valid:"required~Your age is required"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
