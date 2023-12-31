package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"fmt"

	"github.com/oryx-systems/smartduka/pkg/smartduka/domain"
	"github.com/oryx-systems/smartduka/pkg/smartduka/presentation/graph/generated"
)

// SearchUser is the resolver for the searchUser field.
func (r *queryResolver) SearchUser(ctx context.Context, searchTerm string) ([]*domain.User, error) {
	panic(fmt.Errorf("not implemented: SearchUser - searchUser"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) FreezeUser(ctx context.Context, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented: FreezeUser - freezeUser"))
}
func (r *mutationResolver) UnfreezeUser(ctx context.Context, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented: UnfreezeUser - unfreezeUser"))
}
