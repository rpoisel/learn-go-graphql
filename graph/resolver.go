package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/rpoisel/learn-go-graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos map[string]*model.Todo
	users map[string]*model.User
}

func NewResolver() *Resolver {
	return &Resolver{
		todos: make(map[string]*model.Todo),
		users: make(map[string]*model.User),
	}
}
