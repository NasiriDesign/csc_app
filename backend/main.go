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

	//User API
	userRepo := controllers.NewUserRepo()
	r.POST("/adduser", userRepo.CreateUser)
	r.GET("/getusers", userRepo.GetUsers)
	r.GET("/getuser/:user_id", userRepo.GetUser)
	r.PUT("/updateuser/:user_id", userRepo.UpdateUser)
	r.DELETE("/deleteuser/:user_id", userRepo.DeleteUser)

	//Club API
	clubRepo := controllers.NewClubRepo()
	r.POST("/addclub", clubRepo.CreateClub)
	r.GET("/getclubs", clubRepo.GetClubs)
	r.GET("/getclub/:club_id", clubRepo.GetClubByID)
	r.PUT("/updateclub/:club_id", clubRepo.UpdateClubByID)
	r.DELETE("/deleteclub/:club_id", clubRepo.DeleteClub)
	r.GET("/getclubusers/:club_id", clubRepo.GetClubUsers)
	r.POST("/addusertoclub/:club_id/:user_id", clubRepo.AddUserToClub)
	r.POST("/resetuserclub/:user_id", clubRepo.RemoveUserFromClub)

	//Role API
	roleRepo := controllers.NewRoleRepo()
	r.POST("/addrole/:club_id/:user_id", roleRepo.CreateRole)
	r.GET("/getrole/:role_id", roleRepo.GetRoleByID)
	r.PUT("/updaterole", roleRepo.UpdateRole)
	r.POST("/deleterole/:role_id", roleRepo.DeleteRole)

	return r
}
