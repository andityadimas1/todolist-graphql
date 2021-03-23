package api

import (
	"todolist-graphql/config"
	"todolist-graphql/models"
	"todolist-graphql/types"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"get": &graphql.Field{
				Type: types.UserType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(param graphql.ResolveParams) (interface{}, error) {

					ID, ok := param.Args["id"].(int)

					dbPG := config.Connect()
					userVarr := models.User{}
					if ok {
						dbPG.Where("id = ?", ID).First(&userVarr)
					}
					return userVarr, nil
				},
			},
		},
	},
)
