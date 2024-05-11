package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	RoleID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"role_id"`
	ClubID       uuid.UUID // Foreign key
	Name         string    `json:"role_name"`
	EditPost     bool      `json:"edit_post"`
	EditEvents   bool      `json:"edit_events"`
	CancelMember bool      `json:"delete_member"`
	CreatedAt    time.Time `json:"created_at"`
}

// CreateRole creates a new role
func CreateRole(db *gorm.DB, role *Role, clubID, userID uuid.UUID) error {
	// Check if the user ID provided belongs to the club
	var user User
	if err := db.Where("club_id = ? AND user_id = ?", clubID, userID).First(&user).Error; err != nil {
		return errors.New("user not found in the club")
	}

	// Set the club ID and created time
	role.ClubID = clubID
	role.CreatedAt = time.Now()

	// Create the role
	if err := db.Create(role).Error; err != nil {
		return err
	}
	return nil
}

// GetRoleByID retrieves a role by its ID
func GetRoleByID(db *gorm.DB, roleID uuid.UUID) (*Role, error) {
	var role Role
	if err := db.First(&role, "role_id = ?", roleID).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// UpdateRole updates an existing role
func UpdateRole(db *gorm.DB, role *Role) error {
	if err := db.Save(role).Error; err != nil {
		return err
	}
	return nil
}

// DeleteRole deletes an existing role
func DeleteRole(db *gorm.DB, roleID uuid.UUID) error {
	if err := db.Delete(&Role{}, "role_id = ?", roleID).Error; err != nil {
		return err
	}
	return nil
}

func CheckUserPermission(db *gorm.DB, clubID uuid.UUID, userID uuid.UUID, permission string) bool {
	// Retrieve the user's role for the given club
	var role Role
	if err := db.Where("club_id = ? AND user_id = ?", clubID, userID).First(&role).Error; err != nil {
		// If the user doesn't have a role for the club, return false
		return false
	}

	// Check if the role has the specified permission
	switch permission {
	case "edit_post":
		return role.EditPost
	case "edit_events":
		return role.EditEvents
	case "delete_member":
		return role.CancelMember
	default:
		// Handle unknown permission
		return false
	}
}
