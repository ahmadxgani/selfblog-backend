package app

import "go-blog/app/database"

type Account struct {
	Username  string
	Fullname  string
	Image     *string
	CreatedAt string
}

var err error

func ReadAccount(id uint8) (*Account, error) {
	var acc *Account = new(Account)

	err = app.DB.QueryRow("SELECT username, full_name, image, created_at FROM accounts WHERE id=$1", id).Scan(&acc.Username, &acc.Fullname, &acc.Image, &acc.CreatedAt)
	if err != nil {
		return &Account{}, err
	}

	return acc, nil
}
