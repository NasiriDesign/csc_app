package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"user_id"`
	FirstName      string    `gorm:"column:first_name" json:"first_name"`
	LastName       string    `gorm:"column:last_name" json:"last_name"`
	Password       string    `gorm:"column:password" json:"password"`
	DateOfBirth    string    `gorm:"column:date_of_birth;type:DATE" json:"date_of_birth"`
	Address        string    `json:"address"`
	Email          string    `json:"email" binding:"required,email"`
	Phone          string    `json:"phone" binding:"required,e164"`
	Image          []byte    `gorm:"column:image"`
	SubscriptionID int       `gorm:"column:subscription_id" json:"subscription_id"`
	LastLogin      time.Time `gorm:"column:last_login" json:"last_login"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
}

// create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	User.CreatedAt = time.Now()
	User.LastLogin = time.Now()

	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// get users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// get user by id
func GetUser(db *gorm.DB, User *User, id int) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

// delete user
func DeleteUser(db *gorm.DB, User *User, id int) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
