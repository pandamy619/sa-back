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

func (r *queryResolver) PercentGroups(ctx context.Context) ([]*model.PercentGroup, error) {
	pgA := model.PercentGroup{Name: "A", Percent: 80}
	pgB := model.PercentGroup{Name: "B", Percent: 15}
	pgC := model.PercentGroup{Name: "C", Percent: 5}
	pg := []*model.PercentGroup{&pgA, &pgB, &pgC}
	return pg, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
