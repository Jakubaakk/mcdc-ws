package healthcheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthcheckRegister(router *gin.RouterGroup) {
	router.GET("/", GetHealthCheck)
}

func GetHealthCheck(ctx *gin.Context) {
	serializer := HealthcheckSerializer{ctx}
	ctx.JSON(http.StatusOK, serializer.Response(HEALTHY))
}
