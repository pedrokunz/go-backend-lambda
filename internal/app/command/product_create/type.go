package product_create

import (
	"github.com/pedrokunz/go-backend-lambda/internal/app/command/repository/product_create"
	"github.com/pedrokunz/go-backend-lambda/internal/domain/product"
)

type useCase struct {
	repository product_create.Repository
}

type Input struct {
	Name        string
	Description string
	CategoryID  string
}

type Output struct {
	Product *product.Product
}
