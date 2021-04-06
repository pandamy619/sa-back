package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sa-back/graph/generated"
	"sa-back/graph/model"
)

func (r *queryResolver) Fields(ctx context.Context) ([]*model.Field, error) {
	var Fields []*model.Field
	for i := 0; i < 5; i++ {
		dummyField := model.Field{
			Name: fmt.Sprintf("Column name %d", i),
		}
		Fields = append(Fields, &dummyField)
	}

	return Fields, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
