package schema

import (
	"fmt"
	"todolist-graphql/api"

	"github.com/graphql-go/graphql"
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    api.QueryType,    // query itu hanya untuk get data
		Mutation: api.MutationType, //ini untuk Create Update Delete
	},
)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	// check error condition saat run var result
	if len(result.Errors) > 0 {
		fmt.Println("ada error : ", result.Errors)
	}

	return result
}
