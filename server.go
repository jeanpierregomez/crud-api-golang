package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/config"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.Run()
}
