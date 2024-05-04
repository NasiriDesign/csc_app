package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nasiridesign/csc_app/controllers"
)

func main() {
	r := setupRouter()
	_ = r.Run(":4000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/helloworld", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"nachricht": "helloworld",
		})
	})

	userRepo := controllers.New()
	r.POST("/adduser", userRepo.CreateUser)
	r.GET("/getuser", userRepo.GetUsers)
	r.GET("/getuser/:id", userRepo.GetUser)
	r.PUT("/updateuser/:id", userRepo.UpdateUser)
	r.DELETE("/deleteuser/:id", userRepo.DeleteUser)

	return r
}
