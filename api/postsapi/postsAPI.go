package postsapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/helipoc/goapi/database"
	"github.com/helipoc/goapi/models"
	"gorm.io/gorm"
)

func GetAllPosts(c *gin.Context) {
	p := []models.Post{}
	database.DB.Preload("Upvotes").Preload("Downvotes").Find(&p)
	c.JSON(200, p)
}

func DeletePost(c *gin.Context) {
	database.DB.Where("id = ? ", c.Param("id")).Delete(&models.Post{})
	c.JSON(200, gin.H{"success": true})
}

func Upvote(c *gin.Context) {
	var dataReader struct {
		Postid string `json:"postid"`
		Userid string `json:"userid"`
	}
	var res int
	c.BindJSON(&dataReader)

	var p models.Post
	var u models.User

	database.DB.Where("id = ?", dataReader.Postid).Find(&p)
	database.DB.Where("username = ?", dataReader.Userid).Find(&u)
	database.DB.Raw("SELECT user_id FROM user_upvoted WHERE user_id = ? AND post_id = ?;", u.ID, dataReader.Postid).Scan(&res)
	if u.Username == "" || p.Opname == "" {
		fmt.Print("User or Post doesn't exist")
		return
	}
	if res != 0 {
		c.JSON(200, gin.H{"success": false, "msg": "Already upvoted !"})
		return
	}

	database.DB.Model(&p).Association("Upvotes").Append(&u)
	database.DB.Model(&u).Association("Upvoted").Append(&p)
	database.DB.Model(models.Post{}).Where("id = ?", dataReader.Postid).UpdateColumn("score", gorm.Expr("score + ?", 1))
	c.JSON(200, gin.H{"success": true})
	//TODO
}

func Downvote(c *gin.Context) {
	var dataReader struct {
		Postid string `json:"postid"`
		Userid string `json:"userid"`
	}
	var res int

	c.BindJSON(&dataReader)

	var p models.Post
	var u models.User
	database.DB.Where("id = ?", dataReader.Postid).Find(&p)
	database.DB.Where("username = ?", dataReader.Userid).Find(&u)
	if u.Username == "" || p.Opname == "" {
		fmt.Print("User or Post doesn't exist")
		return
	}
	database.DB.Raw("SELECT user_id FROM user_downvoted WHERE user_id = ? AND post_id = ?;", u.ID, dataReader.Postid).Scan(&res)
	if res != 0 {
		c.JSON(200, gin.H{"success": false, "msg": "Already downvoted !"})
		return
	}
	database.DB.Model(&p).Association("Downvotes").Append(&u)
	database.DB.Model(&u).Association("Downvoted").Append(&p)
	database.DB.Model(models.Post{}).Where("id = ?", dataReader.Postid).UpdateColumn("score", gorm.Expr("score - ?", 1))

	c.JSON(200, gin.H{"success": true})
	//TODO
}

func Addpost(c *gin.Context) {
	var md models.Post
	c.BindJSON(&md)
	err := database.DB.Create(&md).Error

	if err != nil {
		c.JSON(400, gin.H{"success": false})
		return
	}

	c.String(200, "Post Ok")

}

func GetUserPosts(c *gin.Context) {
	var posts []models.Post
	err := database.DB.Where("Opname = ?", c.Param("user")).Omit("opname").Find(&posts).Error

	if err != nil {
		c.JSON(400, gin.H{"success": false})
		return
	}

	c.JSON(200, posts)

}

func MountRoutes(p *gin.RouterGroup) {
	p.GET("/", GetAllPosts)
	p.DELETE("/:id", DeletePost)
	p.GET("/:user", GetUserPosts)
	p.POST("/upvote", Upvote)
	p.POST("/downvote", Downvote)
	p.POST("/", Addpost)
}
