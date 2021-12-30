package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"main/graph/model"
)

func (r *mutationResolver) CreateList(ctx context.Context, input *model.NewList) (*model.List, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateList(ctx context.Context, input *model.NewList) (*model.List, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteList(ctx context.Context, id string) (*model.List, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetList(ctx context.Context, id *string) (*model.List, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetListByUser(ctx context.Context, userID *string) ([]*model.List, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTaskByList(ctx context.Context, listID *string) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
