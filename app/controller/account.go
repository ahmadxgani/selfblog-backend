package app

import (
	model "go-blog/app/database/model"
)

func (r *Resolver) Account() (*model.Account, error) {
	acc, err := model.ReadAccount(1)
	if err != nil {
		return &model.Account{}, err
	}
	return acc, nil
}
