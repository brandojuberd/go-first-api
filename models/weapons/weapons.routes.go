package weapons

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(rg *gin.RouterGroup) {
	rg.GET("/weapons", listWeapon)
	rg.GET("/weapon/:id", findWeaponById)
	rg.POST("/weapon", createWeapon)
	rg.POST("/weapon/:id", updateWeapon)
}

func listWeapon(context *gin.Context) {
	fmt.Println(context)
	context.JSON(http.StatusOK, GetList())
}

type UpsertWeaponInput struct {
	Name       string `json:"name"`
	Ammunition int    `json:"ammunition"`
}

func createWeapon(context *gin.Context) {

	var input UpsertWeaponInput
	context.BindJSON(&input)
	context.JSON(http.StatusOK, Create(input))
}

func findWeaponById(context *gin.Context) {
	weaponId := context.Param("id")
	context.JSON(http.StatusOK, FindById(weaponId))
}

func updateWeapon(context *gin.Context) {
	weaponId := context.Param("id")
	var input UpsertWeaponInput
	context.BindJSON(&input)
	context.JSON(http.StatusOK, Update(weaponId, input))
}
