package types

import "github.com/graphql-go/graphql"

var TaskType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Task",
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
