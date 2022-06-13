package product

const (
	sqlProductTable = "product"
)

type SqlProduct struct {
	ID          int     `db:"id"`
	Name        string  `db:"product_name"`
	Description string  `db:"product_description"`
	Unit        string  `db:"unit"`
	Price       float64 `db:"price"`
	TypeId      int     `db:"product_type_id"`
	DiscountId  int     `db:"discount_type_id"`
	TaxId       int     `db:"tax_id"`
}
