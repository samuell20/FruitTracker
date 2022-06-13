package user

import (
	"context"

	"github.com/samuell20/FruitTracker/internal/model"
)

type UserQuery struct {
	repo model.UserRepository
}

func NewUserQuery(repo model.UserRepository) UserQuery {
	return UserQuery{
		repo: repo,
	}
}

func (pq UserQuery) GetAllUsers() ([]model.User, error) {
	return pq.repo.GetAll()
}

func (pq UserQuery) GetUserById(id int, context context.Context) (model.User, error) {
	return pq.repo.Get(id, context)
}
