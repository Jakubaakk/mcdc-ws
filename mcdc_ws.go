package main

import (
	"github.com/gin-gonic/gin"
	"mcdc-ws/configuration"
	"mcdc-ws/healthcheck"
	"mcdc-ws/users"
)

func main() {
	config := configuration.ReadConfiguration()
	db := ConnectToDB(&config)
	r := gin.Default()
	api := r.Group("/api")
	healthcheckRouter := healthcheck.Router(db)
	healthcheckRouter.HealthcheckRegister(api.Group("/health"))
	usersRouter := users.Router(db, &config)
	usersRouter.UsersRegister(api.Group("/v1/users"))
	_ = r.Run()
}
