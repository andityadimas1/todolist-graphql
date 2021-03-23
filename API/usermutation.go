package api

import (
	"todolist-graphql/config"
	"todolist-graphql/models"
	"todolist-graphql/resolver"
	"todolist-graphql/types"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{

			"create": &graphql.Field{
				Type: types.UserType,
				Args: graphql.FieldConfigArgument{

					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"fullname": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"role": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					users := models.User{
						Email:    p.Args["email"].(string),
						Password: p.Args["password"].(string),
						Name:     p.Args["fullName"].(string),
						Role:     p.Args["role"].(string),
					}
					dbPG := config.Connect()

					dbPG.Create(&users)
					return users, nil
				},
			},

			"update": &graphql.Field{
				Type: types.UserType,
				Args: graphql.FieldConfigArgument{ // untuk param
					"id": &graphql.ArgumentConfig{ // id nggak boleh kosong
						Type: graphql.NewNonNull(graphql.Int),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"fullname": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"role": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				// kalau di rest resolve itu kayak controller
				Resolve: resolver.UpdateUserResolve,
			},

			"delete": &graphql.Field{
				Type: types.UserType,
				Args: graphql.FieldConfigArgument{ // untuk param
					"id": &graphql.ArgumentConfig{ // id nggak boleh kosong
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				// kalau di rest resolve itu kayak controller
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// ambil id nya terlebih dahulu, masukan ke dalam sebuah variable
					ID, CheckId := p.Args["id"].(int)

					dbPG := config.Connect()
					userVar := models.User{}

					if CheckId {
						dbPG.Where("id = ?", ID).Delete(&userVar)
					}
					// dbPG.WheDelete(&userVar)

					// dbPG.Save(&userVar)

					// for i, v := range models.User { // product adalah database
					// 	if id == int(v.ID) { // kalau data id dari database == data id dari args
					// 		productVar = models.ProductData[i]
					// 		models.ProductData = append(models.ProductData[:i], models.ProductData[i+1:]...)
					// 	}
					// }
					return userVar, nil // untuk response yang akan ditampilkan
				},
			},
		},
	},
	// graphql.ObjectConfig{
	// 	Name: "task",
	// 	"createTask": &graphql.Field{
	// 		Type: types.TaskType,
	// 		Args: graphql.FieldConfigArgument{

	// 			"tasknama": &graphql.ArgumentConfig{
	// 				Type: graphql.String,
	// 			},
	// 			"completed": &graphql.ArgumentConfig{
	// 				Type: graphql.String,
	// 			},
	// 		},
	// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// 			task := models.Task{
	// 				TaskNama:  p.Args["tasknama"].(string),
	// 				Completed: p.Args["completed"].(string),
	// 			}
	// 			dbPG := config.Connect()

	// 			dbPG.Create(&task)
	// 			return task, nil
	// 		},
	// 	},
	// },
)
