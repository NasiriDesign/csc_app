package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/nasiridesign/csc_app/models"

	"github.com/nasiridesign/csc_app/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// *gormDB repräsentiert eine Verbindung zur Datenbank
type UserRepo struct {
	Db *gorm.DB
}

// Eine Neue Instanz des UserRepo Typs zu erstellen
func NewUserRepo() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

// create user
func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"head": "error", "body": "invalid request"})
		log.Println("INVALID REQUEST:", err)
		return
	}

	// Benutzer in der Datenbank erstellen
	if err := models.CreateUser(repository.Db, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Erfolgreiche Antwort senden
	c.JSON(http.StatusOK, user)
}

// get users
func (repository *UserRepo) GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUser retrieves a user by their UUID.
func (repository *UserRepo) GetUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	var user models.User
	err = models.GetUserByID(repository.Db, &user, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser updates a user's information.
func (repository *UserRepo) UpdateUser(c *gin.Context) {
	var user models.User
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	err = models.GetUserByID(repository.Db, &user, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by their UUID.
func (repository *UserRepo) DeleteUser(c *gin.Context) {
	var user models.User
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	err = models.DeleteUser(repository.Db, &user, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
