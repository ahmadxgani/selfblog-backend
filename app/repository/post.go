package app

import (
	app "go-blog/app/database"
)

type Post struct {
	Title string
  Content string
  Slug string
  Draft bool
}

type UpdatePostData struct {
	Title *string
  Content *string
  Slug *string
  Draft *bool
}

type SqlPost struct {
  Id int32
	Title string
  Content string
  Slug string
  Draft bool
  Likes int32
  Created_At string
  Updated_At string
}

func CreatePost(args Post) (int32, error) {
	_, err := app.DB.Exec("INSERT INTO posts (title, content, slug, draft) VALUES ($1, $2, $3, $4)", args.Title, args.Content, args.Slug, args.Draft)
  if err != nil {
    return 0, err
  }

  return 1, nil
}

func UpdatePost(args UpdatePostData, id int32) (int32, error) {
	_, err := app.DB.Exec("UPDATE posts SET title = COALESCE($1, title), content = COALESCE($2, content), slug = COALESCE($3, slug), draft = COALESCE($4, draft) WHERE id = $5", args.Title, args.Content, args.Slug, args.Draft, id)

  if err != nil {
    return 0, err
  }

  return 1, nil
}

func GetAllPost() ([]*SqlPost, error) {
  var posts []*SqlPost
  rows, err := app.DB.Query("SELECT * FROM posts")
  
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  for rows.Next() {
    var post SqlPost
    err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Slug, &post.Draft, &post.Likes, &post.Created_At, &post.Updated_At)
    if err != nil {
      return nil, err
    }
    posts = append(posts, &post)
  }

  return posts, nil
}

func FindPostByID(id int32) (*SqlPost, error) {
  var post SqlPost
  err := app.DB.QueryRow("SELECT * FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Title, &post.Content, &post.Slug, &post.Draft, &post.Likes, &post.Created_At, &post.Updated_At)
  
  if err != nil {
    return nil, err
  }

  return &post, nil
}

func DeletePost(id int32) (int32, error) {
  _, err := app.DB.Exec("DELETE FROM posts WHERE id = $1", id)
  if err != nil {
    return 0, err
  }
  return 1, nil
}

