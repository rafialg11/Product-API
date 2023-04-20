package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	InputProduct
	UserID uint
	User   *User
}

type InputProduct struct {
	Title       string `json:"title" form:"title" valid:"required~Title of your product is Required"`
	Description string `json:"description" form:"description" valid:"required~Descriptio of your product is Required"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCode := govalidator.ValidateStruct(p)

	if errCode != nil {
		err = errCode
		return
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
