package main

import (
	"net/http"
	"todolist-graphql/config"
	"todolist-graphql/middleware"
	"todolist-graphql/migrator"
	"todolist-graphql/seeder"

	"todolist-graphql/schema"

	"github.com/gin-gonic/gin"
)

func main() {
	dbPG := config.Connect()
	StrDB := middleware.StrDB{DB: dbPG}
	migrator.Migrations(dbPG)
	seeder.TaskSeeder(dbPG)
	seeder.SeederUser(dbPG)

	route := gin.Default()
	route.POST("login", StrDB.MiddleWare().LoginHandler)
	route.POST("/", func(c *gin.Context) {
		type Query struct {
			Query string `json:"query"`
		}

		var query Query

		c.Bind(&query)
		result := schema.ExecuteQuery(query.Query, schema.Schema)
		c.JSON(http.StatusOK, result)
	})

	route.Run()
}
