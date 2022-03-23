// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package user

// Define an input type for the mutation below.
// https://graphql.org/learn/schema/#input-types
//
// Note that, this type is mapped to the generated
// input type in mutation_input.go.
type CreateUserInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Define an input type for the mutation below.
// https://graphql.org/learn/schema/#input-types
//
// Note that, this type is mapped to the generated
// input type in mutation_input.go.
type FollowUserInput struct {
	UserID       int `json:"userID"`
	FollowUserID int `json:"followUserID"`
}