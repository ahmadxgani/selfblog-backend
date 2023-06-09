package app

type Resolver struct{}
type MutationResolver struct{}
type QueryResolver struct{}

func (*Resolver) Mutation() *MutationResolver {
	return &MutationResolver{}
}

func (*Resolver) Query() *QueryResolver {
	return &QueryResolver{}
}
