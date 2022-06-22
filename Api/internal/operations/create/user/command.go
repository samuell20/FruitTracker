package user

import (
	"context"
	"errors"

	"github.com/samuell20/FruitTracker/kit/command"
)

const UserCommandType command.Type = "command.create.User"

// UserCommand is the command dispatched to create a new User.
type UserCommand struct {
	id         int
	username   string
	email      string
	password   string
	role       string
	company_id int
	vat_id     int
}

// NewUserCommand creates a new UserCommand.
func NewUserCommand(id int, username string, companyId int, email string, password string, role string, vatId int) UserCommand {
	return UserCommand{
		id:         id,
		username:   username,
		email:      email,
		password:   password,
		role:       role,
		company_id: companyId,
		vat_id:     vatId,
	}
}

func (c UserCommand) Type() command.Type {
	return UserCommandType
}

// UserCommandHandler is the command handler
// responsible for creating Users.
type UserCommandHandler struct {
	service UserService
}

// NewUserCommandHandler initializes a new UserCommandHandler.
func NewUserCommandHandler(service UserService) UserCommandHandler {
	return UserCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h UserCommandHandler) Handle(ctx context.Context, cmd command.Command) error {

	createUserCmd, ok := cmd.(UserCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateUser(
		ctx,
		createUserCmd.id,
		createUserCmd.username,
		createUserCmd.company_id,
		createUserCmd.email,
		createUserCmd.password,
		createUserCmd.role,
		createUserCmd.vat_id,
	)
}
