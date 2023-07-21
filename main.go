package main

import (
	"time"

	"github.com/gin-gonic/gin"
	route "github.com/jordanlanch/bia-test/api/route"
	"github.com/jordanlanch/bia-test/bootstrap"
)

func main() {

	app := bootstrap.App(".env")

	env := app.Env

	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, app.Postgresql, gin)

	gin.Run(env.ServerAddress)
}
