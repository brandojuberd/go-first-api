package users

import (
	"go-first-api/models"

	"github.com/gin-gonic/gin"
)

func Init(rg *gin.RouterGroup) {
	models.DB.AutoMigrate(&User{})
	Routes(rg)
}
