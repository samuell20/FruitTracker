package user

import (
	"context"

	"github.com/samuell20/FruitTracker/internal/model"
	"github.com/samuell20/FruitTracker/kit/event"
)

// UserService is the default UserService interface
// implementation returned by create.NewUserService.
type UserService struct {
	UserRepository model.UserRepository
	eventBus       event.Bus
}

// NewUserService returns the default Service interface implementation.
func NewUserService(UserRepository model.UserRepository, eventBus event.Bus) UserService {
	return UserService{
		UserRepository: UserRepository,
		eventBus:       eventBus,
	}
}

// CreateUser implements the create.UserService interface.
func (s UserService) CreateUser(ctx context.Context, id int, username string, companyId int, email string, password string, role string, vatId int) error {
	User, err := model.NewUser(id, username, companyId, email, password, role, vatId)
	if err != nil {
		return err
	}

	if err := s.UserRepository.Save(ctx, User); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, User.PullEvents())

}
