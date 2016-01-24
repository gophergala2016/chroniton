package models

import (
	"github.com/gophergala2016/chroniton/utils"
)

func init() {
	utils.ORM.AutoMigrate(&User{})
	utils.ORM.AutoMigrate(&Heartbeat{})
	utils.ORM.AutoMigrate(&Project{})
}
