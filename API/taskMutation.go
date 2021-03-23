package api

import (
	"todolist-graphql/config"
	"todolist-graphql/models"
	"todolist-graphql/types"

	"github.com/graphql-go/graphql"
)

var MutationTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "task",
		Fields: graphql.Fields{

			"create": &graphql.Field{
				Type: types.TaskType,
				Args: graphql.FieldConfigArgument{

					"tasknama": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"completed": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					task := models.Task{
						TaskNama:  p.Args["tasknama"].(string),
						Completed: p.Args["completed"].(string),
					}
					dbPG := config.Connect()

					dbPG.Create(&task)
					return task, nil
				},
			},
			"Delete": &graphql.Field{
				Type: types.TaskType,
				Args: graphql.FieldConfigArgument{ // untuk param
					"id": &graphql.ArgumentConfig{ // id nggak boleh kosong
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				// kalau di rest resolve itu kayak controller
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					ID, CheckId := p.Args["id"].(int)

					dbPG := config.Connect()
					TaskVar := models.Task{}

					if CheckId {
						dbPG.Where("id = ?", ID).Delete(&TaskVar)
					}

					return TaskVar, nil // untuk response yang akan ditampilkan
				},
			},
		},
	},
)
