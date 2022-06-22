package product

import (
	"context"

	"github.com/samuell20/FruitTracker/internal/model"
	"github.com/samuell20/FruitTracker/kit/event"
)

// ProductService is the default ProductService interface
// implementation returned by create.NewProductService.
type DeleteProductService struct {
	ProductRepository model.ProductRepository
	eventBus          event.Bus
}

// NewProductService returns the default Service interface implementation.
func NewDeleteProductService(ProductRepository model.ProductRepository, eventBus event.Bus) DeleteProductService {
	return DeleteProductService{
		ProductRepository: ProductRepository,
		eventBus:          eventBus,
	}
}

// CreateProduct implements the create.ProductService interface.
func (s DeleteProductService) DeleteProduct(ctx context.Context, id int) error {

	if err := s.ProductRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
