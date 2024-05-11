package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Club struct {
	ClubID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"club_id"`
	Admin     uuid.UUID
	Email     string    `json:"email" binding:"required,email"`
	Phone     string    `json:"phone" binding:"required,e164"`
	Roles     []Role    `gorm:"foreignkey:ClubID" json:"roles"` // One-to-many relationship with roles
	Name      string    `json:"club_name"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	Active    bool      `json:"active"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
}

// create a user
func CreateClub(db *gorm.DB, Club *Club) (err error) {
	Club.CreatedAt = time.Now()

	err = db.Create(Club).Error
	if err != nil {
		return err
	}
	return nil
}

// get users
func GetClubs(db *gorm.DB, Club *[]Club) (err error) {
	err = db.Find(Club).Error
	if err != nil {
		return err
	}
	return nil
}

// Get Club by Name
func GetClubName(db *gorm.DB, Club *Club, name string) (err error) {
	err = db.Where("club_name = ?", name).First(Club).Error
	if err != nil {
		return err
	}
	return nil
}

func GetClubByID(db *gorm.DB, Club *Club, id uuid.UUID) (err error) {
	err = db.Where("club_id = ?", id).First(Club).Error
	if err != nil {
		return err
	}
	return nil
}

// Update Club
func UpdateClub(db *gorm.DB, Club *Club) (err error) {
	db.Save(Club)
	return nil
}

// delete user
func DeleteClub(db *gorm.DB, Club *Club, id uuid.UUID) (err error) {
	db.Where("club_id = ?", id).Delete(Club)
	return nil
}

func GetClubUsers(db *gorm.DB, clubID uuid.UUID) ([]User, error) {
	var users []User
	err := db.Where("club_id = ?", clubID).Find(&users).Error
	return users, err
}
