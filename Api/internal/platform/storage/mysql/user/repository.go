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
func (r *UserRepository) Save(ctx context.Context, Product model.User) error {
	/*ProductSQLStruct := sqlbuilder.NewStruct(new(SqlUser))
	query, args := ProductSQLStruct.InsertInto(sqlUserTable, sqlProduct{
		ID:          Product.ID().Value(),
		Name:        Product.Name().String(),
		Description: Product.Description().String(),
		TypeId:      Product.TypeId().Value(),
		Price:       Product.Price().Value(),
		DiscountId:  Product.DiscountId().Value(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("Error trying to persist Product on database: %v", err)
	}
	*/
	return nil
}

func (r *UserRepository) Get(id int, ctx context.Context) (model.User, error) {
	/*ProductSQLStruct := sqlbuilder.NewStruct(new(sqlProduct))
	statement := ProductSQLStruct.SelectFrom(sqlProductTable)
	statement.Where(statement.Equal("id", id))
	query, args := statement.Build()
	rows, _ := r.db.Query(query, args...)
	defer rows.Close()

	var product model.Product
	rows.Scan(ProductSQLStruct.Addr(&product)...)
	log.Println("")
	*/
	return model.User{}, nil
}

//GetAll implements the model.ProductRepository interface.
func (r *UserRepository) GetAll() ([]model.User, error) {
	ProductSQLStruct := sqlbuilder.NewStruct(new(SqlUser))
	statement := ProductSQLStruct.SelectFrom(sqlUserTable)
	query, args := statement.Build()
	rows, _ := r.db.Query(query, args...)
	defer rows.Close()
	userList := []model.User{}
	for rows.Next() {
		var user SqlUser
		err := rows.Scan(ProductSQLStruct.Addr(&user)...)
		scannedUser, err := model.NewUser(
			user.ID,
			user.Username,
			user.CompanyId,
			user.Email,
			user.Password,
			user.VatId)
		if err != nil {
			return nil, fmt.Errorf("Error while getting the users")
		}
		userList = append(userList, scannedUser)
	}

	return userList, nil
}

//Upadate implements the model.ProductRepository interface.
func (r *UserRepository) Update(ctx context.Context, Product model.User) error {
	/*productSQLStruct := sqlbuilder.NewStruct(new(SqlUser))
	query, args := productSQLStruct.Update(sqlUserTable, sqlProduct{
		ID:          Product.ID().Value(),
		Name:        Product.Name().String(),
		Description: Product.Description().String(),
		TypeId:      Product.TypeId().Value(),
		Price:       Product.Price().Value(),
		DiscountId:  Product.DiscountId().Value(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist Product on database: %v", err)
	}
	*/
	return nil
}
