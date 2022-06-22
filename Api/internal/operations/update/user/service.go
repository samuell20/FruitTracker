package user

import (
	"context"

	"github.com/samuell20/FruitTracker/internal/model"
	"github.com/samuell20/FruitTracker/kit/event"
)

// UpdateUserService is the default UpdateUserService interface
// implementation returned by Update.NewUpdateUserService.
type UpdateUserService struct {
	UserRepository model.UserRepository
	eventBus       event.Bus
}

// NewUpdateUserService returns the default Service interface implementation.
func NewUpdateUserService(UserRepository model.UserRepository, eventBus event.Bus) UpdateUserService {
	return UpdateUserService{
		UserRepository: UserRepository,
		eventBus:       eventBus,
	}
}

// UpdateUser implements the Update.UpdateUserService interface.
func (s UpdateUserService) UpdateUser(ctx context.Context, id int, username string, companyId int, email string, password string, role string, vatId int) error {
	User, err := model.NewUser(id, username, companyId, email, password, role, vatId)
	if err != nil {
		return err
	}

	if err := s.UserRepository.Update(ctx, User); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, User.PullEvents())

}
