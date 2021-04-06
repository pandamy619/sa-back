package graph

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"reflect"
	"sa-back/graph/generated"
	"testing"
)

func TestFields(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}})))
	q := `
		query {
  			Fields {
    			name
  			}	
		}
	`
	var resp struct {
		Fields []struct {
			Name string
		}
	}
	c.MustPost(q, &resp)
	if len(resp.Fields) != 5 {
		t.Error("length fields not equal 5")
	}
	colsTest := []string{"Column name 0", "Column name 1", "Column name 2", "Column name 3", "Column name 4"}
	var cols []string
	for _, col := range resp.Fields {
		cols = append(cols, col.Name)
	}
	eq := reflect.DeepEqual(cols, colsTest)
	if !eq {
		t.Error("name cols not equal")
	}
}
