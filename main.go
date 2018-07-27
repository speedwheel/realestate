package main

import (
	"github.com/kataras/iris"
	"github.com/speedwheel/immigrationrealestate/web/bootstrap"
	"github.com/speedwheel/immigrationrealestate/web/routes"
)

func main() {
	app := bootstrap.New()
	app.Bootstrap()
	app.Configure(routes.Configure)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
