package resolver

import (
	"todolist-graphql/config"
	"todolist-graphql/models"

	"github.com/graphql-go/graphql"
)

func UpdateUserResolve(p graphql.ResolveParams) (interface{}, error) {
	// ambil id nya terlebih dahulu, masukan ke dalam sebuah variable
	// id := p.Args["id"].(int)
	Email, checkEmail := p.Args["email"].(string)
	Password, checkPassword := p.Args["password"].(string)
	Name, checkName := p.Args["fullName"].(string)

	// log.Println("ini argsnyaa......", p.Args["name"].(string))
	dbPG := config.Connect()
	userVar := models.User{}
	if checkEmail {
		dbPG.Where("email = ?", Email).First(&userVar)
	}
	if checkName {
		userVar.Name = Name
	}

	if checkPassword {
		userVar.Password = Password
	}

	dbPG.Save(&userVar)

	return userVar, nil // untuk response yang akan ditampilkan
}
