package product

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/samuell20/FruitTracker/internal/model"
)

// ProductRepository is a MySQL model.ProductRepository implementation.
type ProductRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewProductRepository initializes a MySQL-based implementation of model.ProductRepository.
func NewProductRepository(db *sql.DB, dbTimeout time.Duration) *ProductRepository {
	return &ProductRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the model.ProductRepository interface.
func (r *ProductRepository) Save(ctx context.Context, Product model.Product) error {
	ProductSQLStruct := sqlbuilder.NewStruct(new(SqlProduct))
	query, args := ProductSQLStruct.InsertInto(sqlProductTable, SqlProduct{
		Name:        Product.Name().String(),
		Description: Product.Description().String(),
		Unit:        Product.Unit().String(),
		TypeId:      Product.TypeId().Value(),
		Price:       Product.Price().Value(),
		DiscountId:  Product.DiscountId().Value(),
		TaxId:       Product.TaxId().Value(),
	}).Build()
	log.Println(query)
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("Error trying to persist Product on database: %v", err)
	}

	return nil
}

//Get implements the model.ProductRepository interface.
func (r *ProductRepository) Get(id int, ctx context.Context) (model.Product, error) {
	ProductSQLStruct := sqlbuilder.NewStruct(new(SqlProduct))
	statement := ProductSQLStruct.SelectFrom(sqlProductTable)
	statement.Where(statement.Equal("id", id))
	query, args := statement.Build()
	row := r.db.QueryRow(query, args...)

	var product SqlProduct
	row.Scan(ProductSQLStruct.Addr(&product)...)
	productScanned, err := model.NewProduct(
		product.ID,
		product.Name,
		product.Description,
		product.Unit,
		product.Price,
		product.TypeId,
		product.DiscountId,
		product.TaxId)
	if err != nil {
		return model.Product{}, err
	}
	return productScanned, nil
}

//GetAll implements the model.ProductRepository interface.
func (r *ProductRepository) GetAll() ([]model.Product, error) {
	ProductSQLStruct := sqlbuilder.NewStruct(new(SqlProduct))
	statement := ProductSQLStruct.SelectFrom(sqlProductTable)
	query, _ := statement.Build()
	rows, _ := r.db.Query(query)
	defer rows.Close()
	productList := []model.Product{}
	for rows.Next() {
		var product SqlProduct
		err := rows.Scan(ProductSQLStruct.Addr(&product)...)
		productScanned, err := model.NewProduct(product.ID, product.Name, product.Description, product.Unit, product.Price, product.TypeId, product.DiscountId, product.TaxId)
		if err != nil {
			return nil, fmt.Errorf("Error while getting the products")
		}
		productList = append(productList, productScanned)
	}
	return productList, nil
}

//Upadate implements the model.ProductRepository interface.
func (r *ProductRepository) Update(ctx context.Context, Product model.Product) error {
	productSQLStruct := sqlbuilder.NewStruct(new(SqlProduct))
	statement := productSQLStruct.Update(sqlProductTable, SqlProduct{
		ID:          Product.ID().Value(),
		Name:        Product.Name().String(),
		Description: Product.Description().String(),
		TypeId:      Product.TypeId().Value(),
		Price:       Product.Price().Value(),
		DiscountId:  Product.DiscountId().Value(),
	})
	statement.Where(statement.Equal("id", Product.ID().Value()))
	query, args := statement.Build()
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist Product on database: %v", err)
	}

	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, id int) error {
	productSQLStruct := sqlbuilder.NewStruct(new(SqlProduct))
	statement := productSQLStruct.DeleteFrom(sqlProductTable)
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
