package healthcheck

const (
	HEALTHY   = "healthy"
	UNHEALTHY = "unhealthy"
)

type HealthcheckModel struct {
	Status string `json:"status"`
}
