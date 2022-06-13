package user

const (
	sqlUserTable = "user"
)

type SqlUser struct {
	ID        int    `db:"id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CompanyId int    `db:"company_id"`
	VatId     int    `db:"vat_id"`
}
