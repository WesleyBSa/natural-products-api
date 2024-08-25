package services

import (
	"natural-products-api/models"
	"natural-products-api/repositories"
)

func GetProducts() ([]models.Product, error) {
	return repositories.FindAllProducts()
}

func AddProduct(product models.Product) error {
	return repositories.CreateProduct(product)
}

// Outras funcionalidades do serviço podem ser adicionadas aqui
