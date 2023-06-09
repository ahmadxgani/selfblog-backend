package app

import (
	"fmt"
	model "go-blog/app/database/model"
)

func (r *MutationResolver) NewPost(args struct{ Input model.Post }) (int32, error) {
	res, err := model.CreatePost(args.Input)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *MutationResolver) RemovePost(args struct{Id int32}) (int32, error) {
	res, err := model.DeletePost(args.Id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *MutationResolver) UpdatePost(args struct{Input struct{ Id int32; Data model.UpdatePostData }}) (int32, error) {
	if args.Input.Data == (model.UpdatePostData{}) {
		return 0, fmt.Errorf("empty input, nothing was updated")
	}

	res, err := model.UpdatePost(args.Input.Data, args.Input.Id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *QueryResolver) GetPost(args struct{Id int32}) (*model.SqlPost, error) {
	res, err := model.FindPostByID(args.Id)
	
	if err != nil {
		return nil, fmt.Errorf("post not found")
	}

	return res, nil
}

func (r *QueryResolver) ShowAllPost() ([]*model.SqlPost, error) {
	res, err := model.GetAllPost()
	
	if err != nil {
		return nil, err
	}

	return res, nil
}