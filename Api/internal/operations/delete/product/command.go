package product

import (
	"context"
	"errors"

	"github.com/samuell20/FruitTracker/kit/command"
)

const DeleteProductCommandType command.Type = "command.delete.Product"

// DeleteProductCommand is the command dispatched to create a new Product.
type DeleteProductCommand struct {
	id int
}

// NewDeleteProductCommand creates a new DeleteProductCommand.
func NewDeleteProductCommand(id int) DeleteProductCommand {
	return DeleteProductCommand{
		id: id,
	}
}

func (c DeleteProductCommand) Type() command.Type {
	return DeleteProductCommandType
}

// DeleteProductCommandHandler is the command handler
// responsible for creating Products.
type DeleteProductCommandHandler struct {
	service DeleteProductService
}

// NewDeleteProductCommandHandler initializes a new DeleteProductCommandHandler.
func NewDeleteProductCommandHandler(service DeleteProductService) DeleteProductCommandHandler {
	return DeleteProductCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h DeleteProductCommandHandler) Handle(ctx context.Context, cmd command.Command) error {

	deleteProductCmd, ok := cmd.(DeleteProductCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.DeleteProduct(
		ctx,
		deleteProductCmd.id,
	)
}
