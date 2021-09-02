package service

import (
	"belajar-golang-unittest/entity"
	"belajar-golang-unittest/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return category, errors.New("Category not found")
	}

	return category, nil
}
