package controller

type tagResolver struct{}
type postResolver struct{}

// type hello interface {
// 	Hello() string
// }

type accountResolver struct {
	// hello
}
type guestResolver struct{}

func (*accountResolver) Hello() string { return "Hello, world!" }

type Resolver struct{}

func (r *Resolver) Account() (*accountResolver, error) {
	return &accountResolver{}, nil
}
