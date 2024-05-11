package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nasiridesign/csc_app/database"
	"github.com/nasiridesign/csc_app/models"
	"gorm.io/gorm"
)

// *gormDB repräsentiert eine Verbindung zur Datenbank
type ClubRepo struct {
	Db *gorm.DB
}

// Eine Neue Instanz des UserRepo Typs zu erstellen
func NewClubRepo() *ClubRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Club{})
	return &ClubRepo{Db: db}
}

// create user
func (repository *ClubRepo) CreateClub(c *gin.Context) {
	var club models.Club

	if err := c.ShouldBindJSON(&club); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"head": "error", "body": "invalid request"})
		log.Println("INVALID REQUEST:", err)
		return
	}

	// Benutzer in der Datenbank erstellen
	if err := models.CreateClub(repository.Db, &club); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Erfolgreiche Antwort senden
	c.JSON(http.StatusOK, club)
}

// Get Clubs
func (repository *ClubRepo) GetClubs(c *gin.Context) {
	var club []models.Club
	err := models.GetClubs(repository.Db, &club)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, club)
}

// Get Club By ID
func (repository *ClubRepo) GetClubByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("club_id"))
	var club models.Club
	err := models.GetClubByID(repository.Db, &club, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, club)
}

// Update Club
func (repository *ClubRepo) UpdateClubByID(c *gin.Context) {
	var club models.Club
	id, _ := uuid.Parse(c.Param("club_id"))
	err := models.GetClubByID(repository.Db, &club, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&club)
	err = models.UpdateClub(repository.Db, &club)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, club)
}

// Delete Club by UUID
func (repository *ClubRepo) DeleteClub(c *gin.Context) {
	var club models.Club
	id, _ := uuid.Parse(c.Param("club_id"))
	err := models.DeleteClub(repository.Db, &club, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (repository *ClubRepo) GetClubUsers(c *gin.Context) {
	// Parse club_id
	clubID, err := uuid.Parse(c.Param("club_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	// Check if the club exists
	var existingClub models.Club
	if err := models.GetClubByID(repository.Db, &existingClub, clubID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Club not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Fetch users of the club
	users, err := models.GetUsersByClubID(repository.Db, clubID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (repository *ClubRepo) AddUserToClub(c *gin.Context) {
	// Parse user_id
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	// Parse club_id
	clubID, err := uuid.Parse(c.Param("club_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Club ID"})
		return
	}

	// Check if the user exists
	var existingUser models.User
	if err := models.GetUserByID(repository.Db, &existingUser, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Check if the club exists
	var existingClub models.Club
	if err := models.GetClubByID(repository.Db, &existingClub, clubID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Club not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Check if the user already belongs to a club
	if existingUser.ClubID != uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already in a club"})
		return
	}

	// Add user to the club
	if err := models.AddUserToClub(repository.Db, userID, clubID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added to club successfully"})
}

func (repository *ClubRepo) RemoveUserFromClub(c *gin.Context) {
	// Parse user_id
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	// Check if the user exists
	var existingUser models.User
	if err := models.GetUserByID(repository.Db, &existingUser, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Check if the user is already not in a club
	if existingUser.ClubID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not in a club"})
		return
	}

	// Set user's club_id to nil
	existingUser.ClubID = uuid.Nil
	if err := repository.Db.Save(&existingUser).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User removed from club successfully"})
}
