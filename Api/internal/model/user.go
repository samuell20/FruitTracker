package model

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/samuell20/FruitTracker/kit/event"
)

var ErrInvalidUserId = errors.New("Invalid user Id")

type UserId struct {
	value int
}

func NewUserId(id int) (UserId, error) {

	return UserId{
		value: id,
	}, nil
}

func (u UserId) Value() int {
	return u.value
}

var ErrInvalidUserUsername = errors.New("Invalid username")

type UserUsername struct {
	value string
}

func NewUsername(username string) (UserUsername, error) {

	return UserUsername{
		value: username,
	}, nil
}

func (u UserUsername) String() string {
	return u.value
}

var ErrInvalidUserCompanyId = errors.New("Invalid company id")

type UserCompanyId struct {
	value int
}

func NewCompanyId(companyId int) (UserCompanyId, error) {

	return UserCompanyId{value: companyId}, nil
}

func (u UserCompanyId) Value() int {
	return u.value
}

var ErrInvalidUserEmail = errors.New("Invalid email")

type UserEmail struct {
	value string
}

func NewEmail(email string) (UserEmail, error) {

	return UserEmail{value: email}, nil
}

func (u UserEmail) String() string {

	return u.value
}

var ErrInvalidUserPassword = errors.New("Invalid password")

type UserPassword struct {
	value string
}

func NewPassword(password string) (UserPassword, error) {

	return UserPassword{value: password}, nil
}

func (u UserPassword) String() string {
	return u.value
}

var ErrInvalidVatId = errors.New("Invalid user vat_id")

type UserVatId struct {
	value int
}

func NewVatId(id int) (UserVatId, error) {

	return UserVatId{
		value: id,
	}, nil
}

func (u UserVatId) Value() int {
	return u.value
}

type User struct {
	id        UserId
	username  UserUsername
	companyId UserCompanyId
	email     UserEmail
	password  UserPassword
	vatId     UserVatId
	events    []event.Event
}

// ProductRepository defines the expected behaviour from a Product storage.
type UserRepository interface {
	Save(ctx context.Context, User User) error
	GetAll() ([]User, error)
	Get(id int, ctx context.Context) (User, error)
	Update(ctx context.Context, User User) error
}

func NewUser(id int, username string, companyId int, email string, password string, vatId int) (User, error) {

	idVO, err := NewUserId(id)
	if err != nil {
		return User{}, err
	}

	usernameVo, err := NewUsername(username)
	if err != nil {
		return User{}, err
	}
	companyIdVo, err := NewCompanyId(companyId)
	if err != nil {
		return User{}, err
	}
	emailVo, err := NewEmail(email)
	if err != nil {
		return User{}, err
	}
	passwordVo, err := NewPassword(password)
	if err != nil {
		return User{}, err
	}
	vatIdVo, err := NewVatId(vatId)
	if err != nil {
		return User{}, err
	}

	return User{
		id:        idVO,
		username:  usernameVo,
		companyId: companyIdVo,
		email:     emailVo,
		password:  passwordVo,
		vatId:     vatIdVo,
	}, nil
}

func (U User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		CompanyId int    `json:"company_id"`
		VatId     int    `json:"vat_id"`
	}{
		Id:        U.Id().Value(),
		Username:  U.Username().String(),
		Email:     U.Email().String(),
		Password:  U.Password().String(),
		CompanyId: U.CompanyId().Value(),
		VatId:     U.VatId().Value(),
	})
}

func (u User) Id() UserId {
	return u.id
}

func (u User) Username() UserUsername {
	return u.username
}

func (u User) Email() UserEmail {
	return u.email
}

func (u User) Password() UserPassword {
	return u.password
}

func (u User) CompanyId() UserCompanyId {
	return u.companyId
}

func (u User) VatId() UserVatId {
	return u.vatId
}

// Record records a new domain event.
func (u *User) Record(evt event.Event) {
	u.events = append(u.events, evt)
}

// PullEvents returns all the recorded domain events.
func (u User) PullEvents() []event.Event {
	evt := u.events
	u.events = []event.Event{}

	return evt
}
