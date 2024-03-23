package product_create

import (
	"errors"

	"github.com/pedrokunz/go-backend-lambda/internal/app/command/repository/product_create"
	domain "github.com/pedrokunz/go-backend-lambda/internal/domain/product"
)

func New(repository product_create.Repository) (output useCase, err error) {
	if repository == nil {
		return output, errors.New("repository is invalid")
	}

	return useCase{
		repository: repository,
	}, nil
}

func (useCase *useCase) Execute(input Input) (output Output, err error) {
	err = input.Validate()
	if err != nil {
		return output, err
	}

	product := &domain.Product{
		Name:        input.Name,
		Description: input.Description,
		Category: domain.Category{
			ID: input.CategoryID,
		},
	}

	err = useCase.repository.Create(product_create.CreateInput{
		Product: product,
	})
	if err != nil {
		return output, err
	}

	return Output{
		Product: product,
	}, nil
}
