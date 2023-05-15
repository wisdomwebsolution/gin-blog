package tag

import (
	"gin-blog/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var tags []models.Tag

	db, err := models.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	db.Model(&models.Tag{}).Find(&tags)

	c.JSON(http.StatusOK, &tags)
}
