package user

import (
	"context"
	"log"

	"github.com/samuell20/FruitTracker/kit/event"
)

type TestOnUserCreated struct {
}

func NewTestOnUserCreated() TestOnUserCreated {
	return TestOnUserCreated{}
}

func (t TestOnUserCreated) Handle(_ context.Context, evt event.Event) error {

	log.Printf("TestOnUserCreated")
	return nil
}
