package product_create

import "errors"

func (input *Input) Validate() error {
	if input.Name == "" {
		return errors.New("name is required")
	}

	if input.Description == "" {
		return errors.New("description is required")
	}

	if input.CategoryID == "" {
		return errors.New("category ID is required")
	}

	return nil
}
