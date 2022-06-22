package user

import (
	"context"
	"errors"

	"github.com/samuell20/FruitTracker/kit/command"
)

const DeleteUserCommandType command.Type = "command.delete.User"

// DeleteUserCommand is the command dispatched to create a new User.
type DeleteUserCommand struct {
	id int
}

// NewDeleteUserCommand creates a new DeleteUserCommand.
func NewDeleteUserCommand(id int) DeleteUserCommand {
	return DeleteUserCommand{
		id: id,
	}
}

func (c DeleteUserCommand) Type() command.Type {
	return DeleteUserCommandType
}

// DeleteUserCommandHandler is the command handler
// responsible for creating Users.
type DeleteUserCommandHandler struct {
	service DeleteUserService
}

// NewDeleteUserCommandHandler initializes a new DeleteUserCommandHandler.
func NewDeleteUserCommandHandler(service DeleteUserService) DeleteUserCommandHandler {
	return DeleteUserCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h DeleteUserCommandHandler) Handle(ctx context.Context, cmd command.Command) error {

	deleteUserCmd, ok := cmd.(DeleteUserCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.DeleteUser(
		ctx,
		deleteUserCmd.id,
	)
}
