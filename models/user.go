package models

import (
	"product-api/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your Full Name is Required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your Email is Required,email~Invalid Email Format"`
	Password string    `gorm:"not null" json:"-" form:"password" valid:"required~Your password is Required,minstringlength(6)~Password has to have a minimum of 6 characters"`
	Role     string    `gorm:"not null" json:"role" form:"role" valid:"required~Your Role is Required"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
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
