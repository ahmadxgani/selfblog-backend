package app

import (
	"fmt"
	model "go-blog/app/database/model"
)

func (r *QueryResolver) Account() (*model.Account, error) {
	acc, err := model.ReadAccount(1)
	if err != nil {
		return &model.Account{}, err
	}
	return acc, nil
}

func (r *MutationResolver) UpdateAccount(args struct{ Input model.UpdateAccountInput }) (int32, error) {
	if args.Input == (model.UpdateAccountInput{}) {
		return 0, fmt.Errorf("empty input, nothing was updated")
	}

	updated, err := model.UpdateAccount(args.Input)
	if err != nil {
		return 0, err
	}

	return updated, nil
}
