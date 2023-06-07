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
		return nil, err
	}

	return acc, nil
}

type UpdateAccountInput struct {
	Image    *string
	Username *string
	Fullname *string
}

func UpdateAccount(args UpdateAccountInput) (*Account, error) {
	// var acc *Account = new(Account)

	_, err = app.DB.Exec("UPDATE accounts SET username = 'testit' WHERE id=$1", 1)
	updated, err := ReadAccount(1)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
