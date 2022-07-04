package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(rg *gin.RouterGroup) {
	rg.GET("/users", listUsers)
	rg.POST("/user", createUser)
	rg.GET("/user/:id", findUserById)
	rg.POST("/user/:id", updateUser)
}

func listUsers(context *gin.Context){
	context.JSON(http.StatusOK, GetList())
}

func findUserById(context *gin.Context){
	userId := context.Param("id")
	context.JSON(http.StatusOK, FindById(userId))
}

func createUser(context *gin.Context){
	var user User
	context.BindJSON(&user)

	Create(&user)
	context.JSON(http.StatusOK, user)
}

func updateUser(context *gin.Context){
	var user User
	context.BindJSON(&user)

	userId := context.Param("id")
	Update(userId, &user)
	context.JSON(http.StatusOK, user)
}