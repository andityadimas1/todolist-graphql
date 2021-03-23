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
			"getUser": &graphql.Field{
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
			"getTask": &graphql.Field{
				Type: types.TaskType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(param graphql.ResolveParams) (interface{}, error) {
					ID, success := param.Args["id"].(int)

					dbPG := config.Connect()
					taskVar := models.Task{}
					if success {
						dbPG.Where("id = ?", ID).First(taskVar)
					}
					return taskVar, nil
				},
			},
			"ListTask": &graphql.Field{
				Type: graphql.NewList(types.ListTaskType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					dbPG := config.Connect()

					var (
						ListData []models.ListData
					)
					dbPG.Find(ListData)
					return ListData, nil
				},
			},
		},
	},
)
