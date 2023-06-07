package app

import (
	"fmt"
	model "go-blog/app/database/model"
)

func (r *Resolver) Account() (*model.Account, error) {
	acc, err := model.ReadAccount(1)
	if err != nil {
		return &model.Account{}, err
	}
	return acc, nil
}

func (r *Resolver) UpdateAccount(args struct{ Input model.UpdateAccountInput }) (*model.Account, error) {
	if args.Input == (model.UpdateAccountInput{}) {
		return &model.Account{}, fmt.Errorf("empty input, nothing was updated")
	}

	updated, err := model.UpdateAccount(args.Input)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
