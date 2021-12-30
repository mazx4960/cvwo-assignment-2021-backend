package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"main/graph/model"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTask(ctx context.Context, id *string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTaskByUser(ctx context.Context, userID *string) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
