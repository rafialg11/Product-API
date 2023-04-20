package service

import (
	"errors"
	"product-api/models"
	"product-api/repository"
)

type PService interface {
	GetOneProduct(id uint) (*models.Product, error)
	GetAllProduct() (*[]models.Product, error)
}

type ProductService struct {
	Repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *ProductService {
	return &ProductService{repository}
}

func (service *ProductService) GetOneProduct(id uint) (*models.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return product, errors.New("product not found")
	}
	return product, nil
}

func (service *ProductService) GetAllProduct() (*[]models.Product, error) {
	product := service.Repository.FindAll()
	if product == nil || len(*product) == 0 {
		return product, errors.New("data product not available")
	}
	return product, nil
}
