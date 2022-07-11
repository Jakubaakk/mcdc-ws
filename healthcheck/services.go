package healthcheck

import (
	"gorm.io/gorm"
	"mcdc-ws/users"
)

type HealthcheckService struct {
	db *gorm.DB
}

func (self *HealthcheckService) isApplicationHealthy() bool {
	var user users.User
	result := self.db.Model(&users.User{}).First(&user)
	err := result.Error
	if err != nil {
		return false
	}
	return true
}
