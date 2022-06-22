package user

import (
	"context"

	"github.com/samuell20/FruitTracker/internal/model"
	"github.com/samuell20/FruitTracker/kit/event"
)

// DeleteUserService is the default DeleteUserService interface
// implementation returned by create.NewDeleteUserService.
type DeleteUserService struct {
	UserRepository model.UserRepository
	eventBus       event.Bus
}

// NewDeleteUserService returns the default Service interface implementation.
func NewDeleteUserService(UserRepository model.UserRepository, eventBus event.Bus) DeleteUserService {
	return DeleteUserService{
		UserRepository: UserRepository,
		eventBus:       eventBus,
	}
}

// CreateUser implements the create.DeleteUserService interface.
func (s DeleteUserService) DeleteUser(ctx context.Context, id int) error {

	if err := s.UserRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
