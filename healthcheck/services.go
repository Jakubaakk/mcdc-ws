package healthcheck

import (
	"gorm.io/gorm"
	"mcdc-ws/users"
)

type HealthcheckService struct {
	db *gorm.DB
}

func (self *HealthcheckService) isApplicationHealthy() bool {
	var count int64
	result := self.db.Model(&users.User{}).Count(&count)
	err := result.Error
	if err != nil {
		return false
	}
	return true
}
