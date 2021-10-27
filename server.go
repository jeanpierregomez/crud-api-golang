package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/config"
	"github.com/ydhnwb/golang_api/repository"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	r := gin.Default()

	r.Run()
}
