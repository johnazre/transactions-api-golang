package main

import (
	"fmt"

	databaseConfig "github.com/johnazre/transactions-api-gorm/config"

	"github.com/johnazre/transactions-api-gorm/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	types "github.com/johnazre/transactions-api-gorm/models"

	_ "github.com/lib/pq"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	db, _ := databaseConfig.DbStart()

	db.AutoMigrate(&types.User{}, &types.Transaction{})

	fmt.Println("Works")

	app.Get("/users", controllers.GetAllUsers)
	app.Get("/users/{id:int}", controllers.GetOneUser)
	app.Post("/users", controllers.CreateOneUser)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
