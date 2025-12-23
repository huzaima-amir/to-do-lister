package services

import (
	"to-do-lister/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, name, username, password string) (uint, error) {
	user := models.User{
		UserName: username,
		Password: password,
		Name: name,
	}
	if err := db.Create(&user).Error; err != nil{
		return 0, err
	}
	return user.ID, nil
}

func DeleteUserAccount(db *gorm.DB, userID uint) {
	db.Delete(&models.User{}, userID)
}

// finish later
func ChangeName(){

}

func ChangePassword() {

}

func ChangeUsername() {

}