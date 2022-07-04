package users

import (
	"fmt"
	"go-first-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gorm.Model
	ID       uint64 `json:"ID"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"password"`
}
func FindById(userId string)User{
	var user User
	models.DB.Model(&User{}).First(&user, userId)
	return user
}
func GetList() []User {
	var result []User
	models.DB.Model(&User{}).Find(&result)
	fmt.Println(result)
	return result
}

func Create(input *User) {
	models.DB.Model(&User{}).Create(input)
}

func Update(userId string, input *User){
	models.DB.Model(input).Clauses(clause.Returning{}).Where("id", userId).Updates(input)
}

// *** HOOKS ***
func (u *User) AfterFind(tx *gorm.DB) (err error) {
  u.Password = "**hidden**"
  return
}
