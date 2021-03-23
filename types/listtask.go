package types

import "github.com/graphql-go/graphql"

var ListTaskType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ListTask",
		Fields: graphql.Fields{
			"tasknama": &graphql.Field{
				Type: graphql.String,
			},
			"completed": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
