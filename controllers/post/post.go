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
	var postInput NewPost

	if err := c.ShouldBindJSON(&postInput); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	db, err := models.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	var user models.User
	var post models.Post
	db.Where("username =?", username).Find(&user)
	db.Where("id =?", postId).Find(&post)

	if user.ID != uint(post.AuthorID) {
		c.JSON(http.StatusForbidden, "You are not the author")
		return
	}

	db.Model(&post).Where("id =?", postId).Updates(&models.Post{Title: postInput.Title, Content: postInput.Content})

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
	username := c.GetString("username")
	postId := c.Param("id")

	db, err := models.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	var user models.User
	var post models.Post
	db.Where("username =?", username).Find(&user)
	db.Where("id =?", postId).Find(&post)

	if user.ID != uint(post.AuthorID) {
		c.JSON(http.StatusForbidden, "You are not the author")
		return
	}

	db.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
