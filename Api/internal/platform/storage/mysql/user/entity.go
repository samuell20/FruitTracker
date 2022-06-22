package user

const (
	sqlUserTable = "user"
)

type sqlUser struct {
	Id        int    `db:"id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Role      string `db:"role"`
	CompanyId int    `db:"company_id"`
	VatId     int    `db:"vat_id"`
}
