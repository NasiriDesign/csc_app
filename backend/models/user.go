package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"user_id"`
	ClubID         uuid.UUID `json:"club_id"` // Foreign key
	RoleID         uuid.UUID // Foreign key
	Role           Role      `json:"role"`
	FirstName      string    `gorm:"column:first_name" json:"first_name"`
	LastName       string    `gorm:"column:last_name" json:"last_name"`
	Password       string    `json:"password"`
	DateOfBirth    string    `gorm:"column:date_of_birth;type:DATE" json:"date_of_birth"`
	Address        string    `json:"address"`
	Email          string    `json:"email" binding:"required,email"`
	Phone          string    `json:"phone" binding:"required,e164"`
	Image          []byte
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
func GetUserByID(db *gorm.DB, user *User, userID uuid.UUID) error {
	err := db.Where("user_id = ?", userID).First(user).Error
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
func DeleteUser(db *gorm.DB, User *User, id uuid.UUID) (err error) {
	db.Where("user_id = ?", id).Delete(User)
	return nil
}
func AddUserToClub(db *gorm.DB, userID uuid.UUID, clubID uuid.UUID) error {
	// Update the ClubID field of the user with the given UserID
	if err := db.Model(&User{}).Where("user_id = ?", userID).Update("club_id", clubID).Error; err != nil {
		return err
	}
	return nil
}

func GetUsersByClubID(db *gorm.DB, clubID uuid.UUID) ([]User, error) {
	var users []User
	if err := db.Where("club_id = ?", clubID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
