package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/samuell20/FruitTracker/internal/model"
)

// ProductRepository is a MySQL model.ProductRepository implementation.
type UserRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewProductRepository initializes a MySQL-based implementation of model.ProductRepository.
func NewUserRepository(db *sql.DB, dbTimeout time.Duration) *UserRepository {
	return &UserRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the model.ProductRepository interface.
func (r *UserRepository) Save(ctx context.Context, User model.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	query, args := userSQLStruct.InsertInto(sqlUserTable, sqlUser{
		Id:        User.Id().Value(),
		Username:  User.Username().String(),
		Email:     User.Email().String(),
		Password:  User.Password().String(),
		Role:      User.Role().String(),
		CompanyId: User.CompanyId().Value(),
		VatId:     User.VatId().Value(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("Error trying to persist Product on database: %v", err)
	}

	return nil
}

func (r *UserRepository) Get(id int, ctx context.Context) (model.User, error) {
	ProductSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	statement := ProductSQLStruct.SelectFrom(sqlUserTable)
	statement.Where(statement.Equal("id", id))
	query, args := statement.Build()
	row := r.db.QueryRow(query, args...)
	var user sqlUser
	row.Scan(ProductSQLStruct.Addr(&user)...)
	scannedUser, err := model.NewUser(
		user.Id,
		user.Username,
		user.CompanyId,
		user.Email,
		user.Password,
		user.Role,
		user.VatId)
	if err != nil {
		return model.User{}, fmt.Errorf("Error while getting the user")
	}
	return scannedUser, nil
}

//GetAll implements the model.ProductRepository interface.
func (r *UserRepository) GetAll() ([]model.User, error) {
	ProductSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	statement := ProductSQLStruct.SelectFrom(sqlUserTable)
	query, args := statement.Build()
	rows, _ := r.db.Query(query, args...)
	defer rows.Close()
	userList := []model.User{}
	for rows.Next() {
		var user sqlUser
		err := rows.Scan(ProductSQLStruct.Addr(&user)...)
		scannedUser, err := model.NewUser(
			user.Id,
			user.Username,
			user.CompanyId,
			user.Email,
			user.Password,
			user.Role,
			user.VatId)
		if err != nil {
			return nil, fmt.Errorf("Error while getting the users")
		}
		userList = append(userList, scannedUser)
	}

	return userList, nil
}

//Upadate implements the model.ProductRepository interface.
func (r *UserRepository) Update(ctx context.Context, User model.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	statement := userSQLStruct.Update(sqlUserTable, sqlUser{
		Id:        User.Id().Value(),
		Username:  User.Username().String(),
		Email:     User.Email().String(),
		Password:  User.Password().String(),
		Role:      User.Role().String(),
		CompanyId: User.CompanyId().Value(),
		VatId:     User.VatId().Value(),
	})
	statement.Where(statement.Equal("id", User.Id().Value()))
	query, args := statement.Build()
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist Product on database: %v", err)
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	productSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	statement := productSQLStruct.DeleteFrom(sqlUserTable)
	statement.Where(statement.Equal("id", id))
	query, args := statement.Build()
	rows, _ := r.db.Exec(query, args...)
	nRows, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if nRows == 0 {
		return fmt.Errorf("No rows updated")
	}
	return nil
}
