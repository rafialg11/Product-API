package service

import (
	"product-api/models"
	"product-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", uint(1)).Return(nil)

	product, err := productService.GetOneProduct(uint(1))

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "Error response has to be 'product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := models.Product{
		GormModel: models.GormModel{
			ID: 2,
		},
		UserID: 1,
		InputProduct: models.InputProduct{
			Title:       "Bantal",
			Description: "Bantal Nyaman Dipakai",
		},
	}

	productRepository.Mock.On("FindById", uint(2)).Return(product)

	result, err := productService.GetOneProduct(uint(2))

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, product.GormModel.ID, result.GormModel.ID, "Result has to be '2'")
	assert.Equal(t, &product, result, "Result has to be product data with id '2'")
}

func TestProductServiceGetAllProductNotAvailable(t *testing.T) {
	productRepository.Mock.On("FindAll").Return(nil)

	product, err := productService.GetAllProduct()

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "data product not available", err.Error(), "Error response has to be 'data product not available'")
}

func TestProductServiceGetAllProduct(t *testing.T) {
	product := []models.Product{
		{
			GormModel: models.GormModel{
				ID: 1,
			},
			UserID: 2,
			InputProduct: models.InputProduct{
				Title:       "Kasur",
				Description: "Kasur Busa kualitas terbaik",
			},
		},

		{
			GormModel: models.GormModel{
				ID: 2,
			},
			UserID: 2,
			InputProduct: models.InputProduct{
				Title:       "Buku",
				Description: "Buku baru",
			},
		},
	}
	productRepository.Mock.On("FindAll").Return(product)
	result, err := productService.GetAllProduct()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(product), len(*result), "Result lenght")
	assert.Equal(t, product, *result, "Result value")
}
