package model

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/samuell20/FruitTracker/internal/events"
	"github.com/samuell20/FruitTracker/kit/event"
)

// ProductId represents the Product unique identifier.
var ErrInvalidProductId = errors.New("Invalid Product ID")

type ProductId struct {
	value int
}

// NewProductId instantiate the VO for ProductId
func NewProductId(value int) (ProductId, error) {
	return ProductId{
		value: value,
	}, nil
}

// String type converts the ProductId into string.
func (id ProductId) Value() int {
	return id.value
}

// ProductName represents the Product name.
var ErrEmptyProductName = errors.New("The field Product Name can not be empty")

type ProductName struct {
	value string
}

// NewProductName instantiate VO for ProductName
func NewProductName(value string) (ProductName, error) {
	if value == "" {
		return ProductName{}, ErrEmptyProductName
	}

	return ProductName{
		value: value,
	}, nil
}

// String type converts the ProductName into string.
func (name ProductName) String() string {
	return name.value
}

// ProductDescription represents the Product Description.
var ErrEmptyDescription = errors.New("The field Description can not be empty")

type ProductDescription struct {
	value string
}

func NewProductDescription(value string) (ProductDescription, error) {
	if value == "" {
		return ProductDescription{}, ErrEmptyDescription
	}

	return ProductDescription{
		value: value,
	}, nil
}

func (description ProductDescription) String() string {
	return description.value
}

// ProductUnit represents the Product unit.
var ErrEmptyUnit = errors.New("The field unit can not be empty")

type ProductUnit struct {
	value string
}

func NewProductUnit(value string) (ProductUnit, error) {
	if value == "" {
		return ProductUnit{}, ErrEmptyUnit
	}

	return ProductUnit{
		value: value,
	}, nil
}

// String type converts the ProductDescription into string.
func (description ProductUnit) String() string {
	return description.value
}

// ProductId represents the Product unique identifier.
var ErrInvalidProductTypeId = errors.New("Invalid Product type ID")

type ProductTypeId struct {
	value int
}

// NewProductId instantiate the VO for ProductId
func NewProductTypeId(value int) (ProductTypeId, error) {
	return ProductTypeId{
		value: value,
	}, nil
}

// String type converts the ProductId into string.
func (id ProductTypeId) Value() int {
	return id.value
}

// ProductId represents the Product unique identifier.
var ErrInvalidProductPrice = errors.New("Invalid product price ")

type ProductPrice struct {
	value float64
}

// NewProductId instantiate the VO for ProductId
func NewProductPrice(value float64) (ProductPrice, error) {
	return ProductPrice{
		value: value,
	}, nil
}

// String type converts the ProductId into string.
func (id ProductPrice) Value() float64 {
	return id.value
}

// ProductDiscountId represents the Product unique identifier.
var ErrInvalidDiscountId = errors.New("Invalid product discount_id ")

type ProductDiscountId struct {
	value int
}

// NewProductId instantiate the VO for ProductId
func NewProductDiscountId(value int) (ProductDiscountId, error) {
	return ProductDiscountId{
		value: value,
	}, nil
}

// String type converts the ProductId into string.
func (id ProductDiscountId) Value() int {
	return id.value
}

// ProductTaxId represents the Product unique identifier.
var ErrInvalidTaxId = errors.New("Invalid product tax_id ")

type ProductTaxId struct {
	value int
}

// NewProductId instantiate the VO for ProductId
func NewProductTaxId(value int) (ProductTaxId, error) {
	return ProductTaxId{
		value: value,
	}, nil
}

// String type converts the ProductId into string.
func (id ProductTaxId) Value() int {
	return id.value
}

// Product is the data structure that represents a Product.
type Product struct {
	id          ProductId
	name        ProductName
	description ProductDescription
	unit        ProductUnit
	type_id     ProductTypeId
	price       ProductPrice
	discount_id ProductDiscountId
	tax_id      ProductTaxId
	events      []event.Event
}

func (P Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Unit        string  `json:"unit"`
		Price       float64 `json:"price"`
		TypeId      int     `json:"typeId"`
		DiscountId  int     `json:"discountId"`
		TaxId       int     `json:"taxId"`
	}{
		ID:          P.ID().Value(),
		Name:        P.Name().String(),
		Description: P.Description().String(),
		Unit:        P.Unit().String(),
		Price:       P.Price().Value(),
		TypeId:      P.TypeId().Value(),
		DiscountId:  P.DiscountId().Value(),
		TaxId:       P.TaxId().Value(),
	})
}

// ProductRepository defines the expected behaviour from a Product storage.
type ProductRepository interface {
	Save(ctx context.Context, Product Product) error
	GetAll() ([]Product, error)
	Get(id int, ctx context.Context) (Product, error)
	Update(ctx context.Context, Product Product) error
	Delete(ctx context.Context, id int) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ProductRepository

// NewProduct creates a new Product.
func NewProduct(id int, name string, description string, unit string, price float64, type_id int, discount_id int, tax_id int) (Product, error) {
	idVO, err := NewProductId(id)
	if err != nil {
		return Product{}, err
	}

	nameVO, err := NewProductName(name)
	if err != nil {
		return Product{}, err
	}

	descriptionVO, err := NewProductDescription(description)
	if err != nil {
		return Product{}, err
	}

	type_idVO, err := NewProductTypeId(type_id)
	if err != nil {
		return Product{}, err
	}

	priceVO, err := NewProductPrice(price)
	if err != nil {
		return Product{}, err
	}

	unitVO, err := NewProductUnit(unit)
	if err != nil {
		return Product{}, err
	}
	discount_idVO, err := NewProductDiscountId(discount_id)
	if err != nil {
		return Product{}, err
	}
	tax_idVO, err := NewProductTaxId(tax_id)
	if err != nil {
		return Product{}, err
	}

	Product := Product{
		id:          idVO,
		name:        nameVO,
		description: descriptionVO,
		price:       priceVO,
		unit:        unitVO,
		type_id:     type_idVO,
		discount_id: discount_idVO,
		tax_id:      tax_idVO,
	}
	Product.Record(events.NewProductCreatedEvent(idVO.Value(), nameVO.String(), descriptionVO.String(), type_idVO.Value(), priceVO.Value(), discount_idVO.Value()))
	return Product, nil
}

// ID returns the Product unique identifier.
func (c Product) ID() ProductId {
	return c.id
}

// Name returns the Product name.
func (c Product) Name() ProductName {
	return c.name
}

// Description returns the Product description.
func (c Product) Description() ProductDescription {
	return c.description
}

func (c Product) Unit() ProductUnit {
	return c.unit
}

func (c Product) TypeId() ProductTypeId {
	return c.type_id
}

func (c Product) Price() ProductPrice {
	return c.price
}

func (c Product) DiscountId() ProductDiscountId {
	return c.discount_id
}

func (c Product) TaxId() ProductTaxId {
	return c.tax_id
}

// Record records a new domain event.
func (c *Product) Record(evt event.Event) {
	c.events = append(c.events, evt)
}

// PullEvents returns all the recorded domain events.
func (c Product) PullEvents() []event.Event {
	evt := c.events
	c.events = []event.Event{}

	return evt
}
