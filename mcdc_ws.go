package main

import (
	"github.com/gin-gonic/gin"
	"mcdc-ws/healthcheck"
	"mcdc-ws/users"
)

func main() {
	db := ConnectToDB()
	r := gin.Default()
	api := r.Group("/api")
	healthcheckRouter := healthcheck.Router(db)
	healthcheckRouter.HealthcheckRegister(api.Group("/health"))
	usersRouter := users.Router(db)
	usersRouter.UsersRegister(api.Group("/v1/users"))
	_ = r.Run()
}
