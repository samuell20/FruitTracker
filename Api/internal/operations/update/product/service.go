package product

import (
	"context"

	"github.com/samuell20/FruitTracker/internal/model"
	"github.com/samuell20/FruitTracker/kit/event"
)

// UpdateProductService is the default UpdateProductService interface
// implementation returned by create.NewUpdateProductService.
type UpdateProductService struct {
	ProductRepository model.ProductRepository
	eventBus          event.Bus
}

// NewUpdateProductService returns the default Service interface implementation.
func NewUpdateProductService(ProductRepository model.ProductRepository, eventBus event.Bus) UpdateProductService {
	return UpdateProductService{
		ProductRepository: ProductRepository,
		eventBus:          eventBus,
	}
}

// UpdateProduct implements the create.UpdateProductService interface.
func (s UpdateProductService) UpdateProduct(ctx context.Context, id int, name string, description string, unit string, price float64, type_id int, discount_id int, tax_id int) error {
	Product, err := model.NewProduct(id, name, description, unit, price, type_id, discount_id, tax_id)
	if err != nil {
		return err
	}

	if err := s.ProductRepository.Save(ctx, Product); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, Product.PullEvents())

}
