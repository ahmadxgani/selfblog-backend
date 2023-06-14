package app

import (
	app "go-blog/app/database"
)

type Account struct {
	Username  string
	Fullname  string
	Image     *string
	CreatedAt string
}

var err error

func IsValid(email string, password string) bool {
	var is_matched bool
	app.DB.QueryRow("SELECT password = crypt($1, password) AS is_matched FROM accounts WHERE email=$2", password, email).Scan(&is_matched)
	return is_matched
}

func ReadAccount(identifier interface{}) (*Account, error) {
	var acc *Account = new(Account)
	
	err = app.DB.QueryRow("SELECT username, full_name, image, created_at FROM accounts WHERE id=$1", identifier).Scan(&acc.Username, &acc.Fullname, &acc.Image, &acc.CreatedAt)

	if err != nil {
		return nil, err
	}

	return acc, nil
}

type UpdateAccountInput struct {
	Image    *string
	Username *string
	Fullname *string
}

func UpdateAccount(args UpdateAccountInput) (int32, error) {
	// var acc *Account = new(Account)

	_, err = app.DB.Exec("UPDATE accounts SET username = 'testit' WHERE id=$1", 1)
	if err != nil {
		return 0, err
	}

	return 1, nil
}
