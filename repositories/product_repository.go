package repositories

import (
	"natural-products-api/database"
	"natural-products-api/models"
)

func FindAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func CreateProduct(product models.Product) error {
	return database.DB.Create(&product).Error
}

// Outras operações como Update, Delete podem ser adicionadas aqui
