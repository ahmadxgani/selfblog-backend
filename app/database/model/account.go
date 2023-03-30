package app

import (
	"go-blog/app/database"

	"github.com/graph-gophers/graphql-go"
)

type Account struct {
	Username  string
	Fullname  string
	Image     *string
	CreatedAt graphql.Time
}

var err error

func ReadAccount(id uint8) (*Account, error) {
	var acc *Account = new(Account)

	// fmt.Println(app.DB == nil)
	// err = errors.New("test")
	// return &Account{}, err

	// err = app.DB.QueryRow("SELECT username, full_name, image, created_at FROM accounts WHERE id=$1", id).Scan(&acc.Username, &acc.Fullname, sql.NullString{String: *&acc.Image, Valid: true}, &acc.CreatedAt)
	err = app.DB.QueryRow("SELECT username, full_name, image FROM accounts WHERE id=$1", id).Scan(&acc.Username, &acc.Fullname, &acc.Image)
	if err != nil {
		return &Account{}, err
	}

	return acc, nil
}