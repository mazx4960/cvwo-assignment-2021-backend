package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"main/graph/model"
)

func (r *mutationResolver) CreateTag(ctx context.Context, input model.NewTag) (*model.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTag(ctx context.Context, input model.NewTag) (*model.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTag(ctx context.Context, id *string) (*model.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTagByUser(ctx context.Context, userID *string) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTaskByTag(ctx context.Context, tagID *string) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
