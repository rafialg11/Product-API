package repository

import (
	"product-api/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindById(id uint) *models.Product
	FindAll() *[]models.Product
}

type PRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *PRepository {
	return &PRepository{db}
}

func (re *PRepository) FindById(id uint) *models.Product {
	product := models.Product{}
	err := re.db.Debug().Joins("User").First(&product, id).Error
	if err != nil {
		return nil
	}
	return &product
}

func (re *PRepository) FindAll() *[]models.Product {
	product := []models.Product{}
	err := re.db.Debug().Joins("User").Order("id ASC").Find(&product).Error
	if err != nil {
		return nil
	}
	return &product
}
