package app

import (
	"fmt"
	model "go-blog/app/database/model"
)

func (r *QueryResolver) AuthToken(args struct{Input struct{Email string; Password string}}) (string, error) {
	is_matched := model.IsValid(args.Input.Email, args.Input.Password)

	if !is_matched {
		return "", fmt.Errorf("email or password is invalid")
	}

	return "JWT token", nil
}
