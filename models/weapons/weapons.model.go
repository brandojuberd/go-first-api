package weapons

import (
	"fmt"
	"go-first-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Weapon struct {
	gorm.Model
	ID         uint   `json:"ID"`
	Name       string `json:"name"`
	Ammunition int    `json:"ammunition"`
}

func Init(rg *gin.RouterGroup) {
	models.DB.AutoMigrate(&Weapon{})
	Routes(rg)
	// models.DB.Create(&Weapon{Name: "D42", Ammunition: 100})
}

func FindById(weaponId string) Weapon {
	var weapon Weapon
	models.DB.Model(&Weapon{}).First(&weapon, weaponId)
	return weapon
}

func GetList() []Weapon {
	// var weapon Weapon
	var result []Weapon
	model := models.DB.Model(&Weapon{}).Find(&result)
	fmt.Println(model)
	return result
}

func Create(input UpsertWeaponInput) Weapon {
	weapon := Weapon{
		Name:       input.Name,
		Ammunition: input.Ammunition,
	}
	models.DB.Model(&Weapon{}).Create(&weapon)
	return weapon
}

func Update(weaponId string, updateInput UpsertWeaponInput) Weapon {
	var currentWeapon = FindById(weaponId)
	var weapon = Weapon{
		Name:       updateInput.Name,
		Ammunition: updateInput.Ammunition,
	}
	models.DB.Model(&currentWeapon).Where("id", weaponId).Updates(weapon)
	return currentWeapon
}
