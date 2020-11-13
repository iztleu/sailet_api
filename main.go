package main

import (
	"fmt"

	"github.com/iztleu/sailet_api/database/mongodb"
	"github.com/iztleu/sailet_api/models"
)

func main() {
	err := mongodb.InitDatabaseConnection("localhost:27017")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mongodb.CloseDatabaseConnection()
	s := mongodb.GetSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	model := models.NewMgModel(s)

	acc, err := model.GetAccount("d.iztleu@gmail.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*acc)

	// acc := &models.Account{
	// 	Email:    "d.iztleu@gmail.com",
	// 	Password: "123456",
	// 	Info:     "Test",
	// }

	acc.Info = "New Test"

	if ok, err := model.CreateOrUpdateAccount(acc); ok == false || err != nil {
		fmt.Println(err)
		return
	}

}
