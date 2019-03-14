package controllers

import (
	"github.com/johnazre/transactions-api-gorm/config"
	"github.com/johnazre/transactions-api-gorm/models"
	"github.com/kataras/iris"
	"time"
)

func GetAllUsers(ctx iris.Context) {
	// Create connection to database
	db, _ := databaseConfig.DbStart()
	// Close connection when function is done running
	defer db.Close()

	// Create container for all users
	var users []types.User

	// Query database for all users and populate container with said results
	db.Find(&users)

	// Respond to request with JSON of all the users
	ctx.JSON(users)
}

func GetOneUser(ctx iris.Context) {
	// Create connection to database
	db, _ := databaseConfig.DbStart()
	// Close connection when function is done running
	defer db.Close()

	// Create container for one user
	var user types.User

	// Acquire the id via the url params
	urlParam, _ := ctx.Params().GetInt("id")

	// Query database for user with a certain ID
	db.First(&user, urlParam)

	ctx.JSON(user)

}

func CreateOneUser(ctx iris.Context) {
	// Create connection to database
	db, _ := databaseConfig.DbStart()
	// Close connection when function is done running
	defer db.Close()

	var requestBody types.User

	// Access request body
	ctx.ReadJSON(&requestBody)

	// Populate new user instance
	user := types.User{
		Email:     requestBody.Email,
		Password:  requestBody.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Create new user in the database
	db.NewRecord(user)
	db.Create(&user)

	var newUser types.User

	db.Where("email = ?", user.Email).Find(&newUser)

	ctx.JSON(newUser)
}
