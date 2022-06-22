package user

import (
	"context"
	"errors"

	"github.com/samuell20/FruitTracker/kit/command"
)

const UpdateUserCommandType command.Type = "command.update.Product"

// UpdateUserCommand is the command dispatched to create a new Product.
type UpdateUserCommand struct {
	id         int
	username   string
	email      string
	password   string
	role       string
	company_id int
	vat_id     int
}

// NewUpdateUserCommand creates a new UpdateUserCommand.
func NewUpdateUserCommand(id int, username string, email string, password string, role string, companyId int, vatId int) UpdateUserCommand {
	return UpdateUserCommand{
		id:         id,
		username:   username,
		email:      email,
		password:   password,
		role:       role,
		company_id: companyId,
		vat_id:     vatId,
	}
}

func (c UpdateUserCommand) Type() command.Type {
	return UpdateUserCommandType
}

// UpdateUserCommandHandler is the command handler
// responsible for creating Products.
type UpdateUserCommandHandler struct {
	service UpdateUserService
}

// NewUpdateUserCommandHandler initializes a new UpdateUserCommandHandler.
func NewUpdateUserCommandHandler(service UpdateUserService) UpdateUserCommandHandler {
	return UpdateUserCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h UpdateUserCommandHandler) Handle(ctx context.Context, cmd command.Command) error {

	createProductCmd, ok := cmd.(UpdateUserCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.UpdateUser(
		ctx,
		createProductCmd.id,
		createProductCmd.username,
		createProductCmd.company_id,
		createProductCmd.email,
		createProductCmd.password,
		createProductCmd.role,
		createProductCmd.vat_id,
	)
}
