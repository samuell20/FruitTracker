package product

import (
	"context"
	"log"

	"github.com/samuell20/FruitTracker/kit/event"
)

type TestOnProductCreated struct {
}

func NewTestOnProductCreated() TestOnProductCreated {
	return TestOnProductCreated{}
}

func (t TestOnProductCreated) Handle(_ context.Context, evt event.Event) error {

	log.Printf("TestOnProductCreated")
	return nil
}
