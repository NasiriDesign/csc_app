package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	PostID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"role_id"`
	ClubID    uuid.UUID `json:"club_id"` // Foreign key
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// CreatePost creates a new post for a club
func CreatePost(db *gorm.DB, post *Post) error {
	// Set the creation timestamp
	post.CreatedAt = time.Now()

	// Check if the user has the necessary permissions to create a post
	if !CheckUserPermission(db, post.ClubID, post.ClubID, "create_post") {
		return errors.New("permission denied")
	}

	// Create the post
	if err := db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

// GetPostsByClubID retrieves all posts for a club
func GetPostsByClubID(db *gorm.DB, clubID uuid.UUID) ([]Post, error) {
	var posts []Post
	if err := db.Where("club_id = ?", clubID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// UpdatePost updates an existing post
func UpdatePost(db *gorm.DB, post *Post) error {
	// Check if the user has the necessary permissions to update the post
	if !CheckUserPermission(db, post.ClubID, post.ClubID, "update_post") {
		return errors.New("permission denied")
	}

	// Update the post
	if err := db.Save(post).Error; err != nil {
		return err
	}
	return nil
}

// DeletePost deletes an existing post
func DeletePost(db *gorm.DB, postID uuid.UUID, userID uuid.UUID) error {
	var post Post
	if err := db.Where("post_id = ? AND user_id = ?", postID, userID).First(&post).Error; err != nil {
		return err
	}

	// Check if the user has the necessary permissions to delete the post
	if !CheckUserPermission(db, post.ClubID, userID, "delete_post") {
		return errors.New("permission denied")
	}

	// Delete the post
	if err := db.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
