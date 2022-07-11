package healthcheck

import "github.com/gin-gonic/gin"

const (
	HEALTHY   = "healthy"
	UNHEALTHY = "unhealthy"
)

type HealthcheckSerializer struct {
	C *gin.Context
}

func (self *HealthcheckSerializer) Response(status string) gin.H {
	return gin.H{
		"status": status,
	}
}
