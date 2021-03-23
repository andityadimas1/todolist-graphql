package main

import (
	"net/http"
	"todolist-graphql/config"
	"todolist-graphql/executions"
	"todolist-graphql/migrator"
	"todolist-graphql/seeder"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/schema"
)

func main() {
	dbPG := config.Connect()
	migrator.Migrations(dbPG)
	seeder.TaskSeeder(dbPG)
	seeder.SeederUser(dbPG)

	route := gin.Default()
	route.POST("/", func(c *gin.Context) {
		type Query struct {
			Query string `json:"query"`
		}

		var query Query

		c.Bind(&query)
		result := executions.ExecuteQuery(query.Query, schema.Schema)

		c.JSON(http.StatusOK, result)
	})

	route.Run()
}
