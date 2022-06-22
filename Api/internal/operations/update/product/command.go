package product

import (
	"context"
	"errors"

	"github.com/samuell20/FruitTracker/kit/command"
)

const UpdateProductCommandType command.Type = "command.update.Product"

// UpdateProductCommand is the command dispatched to update a new Product.
type UpdateProductCommand struct {
	id          int
	name        string
	description string
	unit        string
	price       float64
	type_id     int
	discount_id int
	tax_id      int
}

// NewUpdateProductCommand updates a new UpdateProductCommand.
func NewUpdateProductCommand(id int, name string, description string, unit string, price float64, type_id int, discount_id int, tax_id int) UpdateProductCommand {
	return UpdateProductCommand{
		id:          id,
		name:        name,
		description: description,
		unit:        unit,
		price:       price,
		type_id:     type_id,
		discount_id: discount_id,
		tax_id:      tax_id,
	}
}

func (c UpdateProductCommand) Type() command.Type {
	return UpdateProductCommandType
}

// UpdateProductCommandHandler is the command handler
// responsible for creating Products.
type UpdateProductCommandHandler struct {
	service UpdateProductService
}

// NewUpdateProductCommandHandler initializes a new UpdateProductCommandHandler.
func NewUpdateProductCommandHandler(service UpdateProductService) UpdateProductCommandHandler {
	return UpdateProductCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h UpdateProductCommandHandler) Handle(ctx context.Context, cmd command.Command) error {

	updateProductCmd, ok := cmd.(UpdateProductCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.UpdateProduct(
		ctx,
		updateProductCmd.id,
		updateProductCmd.name,
		updateProductCmd.description,
		updateProductCmd.unit,
		updateProductCmd.price,
		updateProductCmd.type_id,
		updateProductCmd.discount_id,
		updateProductCmd.tax_id,
	)
}
