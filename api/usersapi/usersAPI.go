package usersapi

import (
	"github.com/gin-gonic/gin"
	"github.com/helipoc/goapi/database"
	"github.com/helipoc/goapi/models"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *gin.Context) {
	u := []models.User{}
	database.DB.Preload("Upvoted").Find(&u)
	c.JSON(200, u)
}

func CreateUser(c *gin.Context) {
	var ur models.User
	var dataReader struct {
		Username string
		Password string
	}
	c.BindJSON(&dataReader)

	if dataReader.Username == "" || dataReader.Password == "" {
		c.JSON(400, gin.H{
			"success": false,
			"msg":     "Missing Username AND/OR Password"})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(dataReader.Password), 10)
	if err != nil {
		panic("Error while hashing password")
	}
	ur.Username = dataReader.Username
	ur.Password = string(hashedPass)
	res := database.DB.Create(&ur).Error
	if res != nil {
		c.JSON(400, gin.H{"success": false, "msg": "username already exists !"})
		return
	}

	c.String(200, "OK")

}

func DeleteUser(c *gin.Context) {
	database.DB.Where("id = ? ", c.Param("id")).Delete(&models.User{})
	c.JSON(200, gin.H{"success": true})
}

func MountRoutes(p *gin.RouterGroup) {
	p.GET("/", GetAllUsers)
	p.DELETE("/:id", DeleteUser)
	p.POST("/", CreateUser)
}
