package post

import (
	"gin-blog/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostShow struct {
	models.Post
	Author models.UserPublic `json:"author"`
}

func Index(c *gin.Context) {
	var posts []PostShow
	db, err := models.Database()

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Model(&models.Post{}).Joins("Author", db.Select("username", "id")).Find(&posts)

	c.JSON(http.StatusOK, posts)
}

type NewPost struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"Content" binding:"required"`
}

func Create(c *gin.Context) {
	username := c.GetString("username")
	var newPostInput NewPost

	if err := c.ShouldBindJSON(&newPostInput); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	db, err := models.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	var user models.User
	db.Where("username = ?", username).Find(&user)

	newPost := models.Post{Title: newPostInput.Title, Content: newPostInput.Content, Author: user}

	db.Create(&newPost)

	c.JSON(http.StatusOK, newPost)
}

func Update(c *gin.Context) {
	username := c.GetString("username")
	postId := c.Param("id")
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	db, err := models.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	var user models.User
	db.Where("username =?", username).Find(&user)

	post.Author = user

	db.Model(&models.Post{}).Where("id =?", postId).Updates(&post)

	c.JSON(http.StatusOK, post)
}

func Show(c *gin.Context) {
	postId := c.Param("id")

	db, err := models.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	var post PostShow
	db.Model(&models.Post{}).Where("posts.id = ?", postId).Joins("Author").First(&post)
	c.JSON(http.StatusOK, &post)
}

func Delete(c *gin.Context) {
	postId := c.Param("id")

	db, err := models.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	db.Where("id =?", postId).Delete(&models.Post{})
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
