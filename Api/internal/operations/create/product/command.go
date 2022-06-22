package product

import (
	"context"
	"errors"
	"log"

	"github.com/samuell20/FruitTracker/kit/command"
)

const ProductCommandType command.Type = "command.create.Product"

// ProductCommand is the command dispatched to create a new Product.
type ProductCommand struct {
	name        string
	description string
	unit        string
	price       float64
	type_id     int
	discount_id int
	tax_id      int
}

// NewProductCommand creates a new ProductCommand.
func NewProductCommand(name string, description string, unit string, price float64, type_id int, discount_id int, tax_id int) ProductCommand {
	return ProductCommand{
		name:        name,
		description: description,
		unit:        unit,
		price:       price,
		type_id:     type_id,
		discount_id: discount_id,
		tax_id:      tax_id,
	}
}

func (c ProductCommand) Type() command.Type {
	return ProductCommandType
}

// ProductCommandHandler is the command handler
// responsible for creating Products.
type ProductCommandHandler struct {
	service ProductService
}

// NewProductCommandHandler initializes a new ProductCommandHandler.
func NewProductCommandHandler(service ProductService) ProductCommandHandler {
	return ProductCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h ProductCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	log.Println(cmd)
	createProductCmd, ok := cmd.(ProductCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateProduct(
		ctx,
		createProductCmd.name,
		createProductCmd.description,
		createProductCmd.unit,
		createProductCmd.price,
		createProductCmd.type_id,
		createProductCmd.discount_id,
		createProductCmd.tax_id,
	)
}
