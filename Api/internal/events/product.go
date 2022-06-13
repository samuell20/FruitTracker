package events

import (
	"github.com/samuell20/FruitTracker/kit/event"
)

const ProductCreatedEventType event.Type = "events.Product.created"

type ProductCreatedEvent struct {
	event.BaseEvent
	id          int
	name        string
	description string
	type_id     int
	price       float64
	discount_id int
}

func NewProductCreatedEvent(id int, name, description string, type_id int, price float64, discount_id int) ProductCreatedEvent {
	return ProductCreatedEvent{
		id:          id,
		name:        name,
		description: description,
		type_id:     type_id,
		price:       price,
		discount_id: discount_id,

		/*BaseEvent: event.NewBaseEvent(id),*/
	}
}

func (e ProductCreatedEvent) Type() event.Type {
	return ProductCreatedEventType
}

func (e ProductCreatedEvent) ProductID() int {
	return e.id
}

func (e ProductCreatedEvent) ProductName() string {
	return e.name
}

func (e ProductCreatedEvent) ProductDescription() string {
	return e.description
}

func (e ProductCreatedEvent) ProductTypeId() int {
	return e.type_id
}

func (e ProductCreatedEvent) ProductPrice() float64 {
	return e.price
}

func (e ProductCreatedEvent) ProductDiscountId() int {
	return e.discount_id
}
