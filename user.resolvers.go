package user

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"user/ent"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input CreateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.Create().SetName(input.Name).SetAge(input.Age).Save(ctx)
}

func (r *mutationResolver) FollowUser(ctx context.Context, input FollowUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.UpdateOneID(input.UserID).AddFollowingIDs(input.FollowUserID).Save(ctx)
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
