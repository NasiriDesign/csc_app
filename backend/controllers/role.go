package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nasiridesign/csc_app/database"
	"github.com/nasiridesign/csc_app/models"
	"gorm.io/gorm"
)

// RoleRepo represents the repository for managing roles
type RoleRepo struct {
	Db *gorm.DB
}

// NewRoleRepo creates a new instance of RoleRepo
func NewRoleRepo() *RoleRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Role{})
	return &RoleRepo{Db: db}
}

// CreateRole creates a new role
func (repository *RoleRepo) CreateRole(c *gin.Context) {
	var role models.Role

	// Bind JSON data to role struct
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		log.Println("Invalid request:", err)
		return
	}

	// Ensure user ID belongs to the club
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Ensure user ID belongs to the club
	clubID, err := uuid.Parse(c.Param("club_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid club ID"})
		return
	}

	// Create the role
	if err := models.CreateRole(repository.Db, &role, clubID, userID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	// Successful response
	c.JSON(http.StatusOK, role)
}

// GetRoleByID retrieves a role by its ID
func (repository *RoleRepo) GetRoleByID(c *gin.Context) {
	roleID, err := uuid.Parse(c.Param("role_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := models.GetRoleByID(repository.Db, roleID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch role"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// UpdateRole updates an existing role
func (repository *RoleRepo) UpdateRole(c *gin.Context) {
	var role models.Role

	// Bind JSON data to role struct
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		log.Println("Invalid request:", err)
		return
	}

	// Update the role
	if err := models.UpdateRole(repository.Db, &role); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
		return
	}

	// Successful response
	c.JSON(http.StatusOK, role)
}

// DeleteRole deletes an existing role
func (repository *RoleRepo) DeleteRole(c *gin.Context) {
	roleID, err := uuid.Parse(c.Param("role_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := models.DeleteRole(repository.Db, roleID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
