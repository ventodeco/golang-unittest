package repository

import "belajar-golang-unittest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
