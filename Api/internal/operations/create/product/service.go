package product

import (
	"context"

	"github.com/samuell20/FruitTracker/internal/model"
	"github.com/samuell20/FruitTracker/kit/event"
)

// ProductService is the default ProductService interface
// implementation returned by create.NewProductService.
type ProductService struct {
	ProductRepository model.ProductRepository
	eventBus          event.Bus
}

// NewProductService returns the default Service interface implementation.
func NewProductService(ProductRepository model.ProductRepository, eventBus event.Bus) ProductService {
	return ProductService{
		ProductRepository: ProductRepository,
		eventBus:          eventBus,
	}
}

// CreateProduct implements the create.ProductService interface.
func (s ProductService) CreateProduct(ctx context.Context, id int, name string, description string, unit string, price float64, type_id int, discount_id int, tax_id int) error {
	Product, err := model.NewProduct(id, name, description, unit, price, type_id, discount_id, tax_id)
	if err != nil {
		return err
	}

	if err := s.ProductRepository.Save(ctx, Product); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, Product.PullEvents())

}
