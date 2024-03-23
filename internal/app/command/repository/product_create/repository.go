package product_create

import "github.com/pedrokunz/go-backend-lambda/internal/domain/product"

type CreateInput struct {
	Product *product.Product
}

type Repository interface {
	Create(input CreateInput) error
}
