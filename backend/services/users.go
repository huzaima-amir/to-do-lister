package services

import (
    "fmt"
    "to-do-lister/models"

    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)


func CreateUser(db *gorm.DB, name, username, password string) (uint, error) {
    var count int64
    db.Model(&models.User{}).Where("user_name = ?", username).Count(&count)
    if count > 0 {  //check if username already exists
        return 0, fmt.Errorf("username already taken") 
    }

    // hash password
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return 0, fmt.Errorf("failed to hash password")
    }
    user := models.User{
        Name:     name,
        UserName: username,
        Password: string(hashed),
    }
    if err := db.Create(&user).Error; err != nil {
        return 0, err
    }
    return user.ID, nil
}


func DeleteUser(db *gorm.DB, userID uint) error { //delete account along with all related tasks, tags and events
	return db.Delete(&models.User{}, userID).Error
}

func ChangeName(db *gorm.DB, userID uint, newName string) error {
    return db.Model(&models.User{}).
        Where("id = ?", userID).
        Update("name", newName).Error
}

func ChangeUsername(db *gorm.DB, userID uint, newUsername string) error {
    // check if username already exists
    var count int64
    db.Model(&models.User{}).Where("user_name = ?", newUsername).Count(&count)
    if count > 0 {
        return fmt.Errorf("username already taken")
    }
    // update
    return db.Model(&models.User{}).
        Where("id = ?", userID).
        Update("user_name", newUsername).Error
}

func ChangePassword(db *gorm.DB, userID uint, oldPassword, newPassword string) error {
    var user models.User

    // load user
    if err := db.First(&user, userID).Error; err != nil {
        return fmt.Errorf("user not found")
    }
    // verify old password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
        return fmt.Errorf("incorrect old password")
    }
    //hash new password
    hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("failed to hash new password")
    }
    // save new password
    return db.Model(&user).Update("password", string(hashed)).Error
}

func ValidateUserCredentials(db *gorm.DB, username, password string) (uint, error) { // helper for loginhandler
    var user models.User

    // Find user by username, in db
    if err := db.Where("user_name = ?", username).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return 0, fmt.Errorf("invalid username or password")
        }
        return 0, err
    }
    //compare hashed password with provided password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return 0, fmt.Errorf("invalid username or password")
    }

    // Return user ID if valid
    return user.ID, nil
}
