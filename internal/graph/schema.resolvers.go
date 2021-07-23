package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/moducate/heimdall/internal/graph/model"

	"github.com/moducate/heimdall/internal/graph/generated"
)

func (r *queryResolver) Schools(ctx context.Context) ([]*model.School, error) {
	schools, err := r.Env.Services.School.GetAll()
	return schools, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
