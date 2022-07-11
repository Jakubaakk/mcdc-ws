package main

import (
	"github.com/gin-gonic/gin"
	"mcdc-ws/healthcheck"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	healthcheck.HealthcheckRegister(api.Group("/health"))
	_ = r.Run()
}
