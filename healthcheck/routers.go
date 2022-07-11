package healthcheck

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type HealthcheckRouter struct {
	service *HealthcheckService
}

func (self *HealthcheckRouter) HealthcheckRegister(router *gin.RouterGroup) {
	router.GET("/", self.getHealthCheck)
}

func Router(db *gorm.DB) HealthcheckRouter {
	return HealthcheckRouter{&HealthcheckService{db}}
}

func (self *HealthcheckRouter) getHealthCheck(ctx *gin.Context) {
	if self.service.isApplicationHealthy() == false {
		ctx.IndentedJSON(http.StatusInternalServerError, HealthcheckModel{UNHEALTHY})
		return
	}
	ctx.IndentedJSON(http.StatusOK, HealthcheckModel{HEALTHY})
}
