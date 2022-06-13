package product

import (
	"context"

	"github.com/samuell20/FruitTracker/internal/model"
)

type ProductQuery struct {
	repo model.ProductRepository
}

func NewProductQuery(repo model.ProductRepository) ProductQuery {
	return ProductQuery{
		repo: repo,
	}
}

func (pq ProductQuery) GetAllProducts() ([]model.Product, error) {
	return pq.repo.GetAll()
}

func (pq ProductQuery) GetProductById(id int, context context.Context) (model.Product, error) {
	return pq.repo.Get(id, context)
}
