package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm/clause"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("UserData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.UserID = userID

	err := db.Debug().Create(&comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("UserData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	comment := []models.Comment{}
	responseData := []map[string]interface{}{}
	err := db.Preload("User").Preload("Photo").Where("user_id=?", userID).Find(&comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for i := range comment {
		nestedUser := map[string]interface{}{
			"id":       comment[i].User.ID,
			"email":    comment[i].User.Email,
			"username": comment[i].User.Username,
		}
		nestedPhoto := map[string]interface{}{
			"id":        comment[i].Photo.ID,
			"title":     comment[i].Photo.Title,
			"caption":   comment[i].Photo.Caption,
			"photo_url": comment[i].Photo.PhotoURL,
			"user_id":   comment[i].Photo.UserID,
		}
		data := map[string]interface{}{
			"id":         comment[i].ID,
			"message":    comment[i].Message,
			"photo_id":   comment[i].PhotoID,
			"user_id":    comment[i].UserID,
			"created_at": comment[i].CreatedAt,
			"updated_at": comment[i].UpdatedAt,
			"User":       nestedUser,
			"Photo":      nestedPhoto,
		}

		responseData = append(responseData, data)
	}

	c.JSON(http.StatusOK, responseData)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	comment := models.Comment{}

	commentID, _ := strconv.Atoi(c.Param("commentID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	err := db.Model(&comment).Where("id=?", commentID).Clauses(clause.Returning{}).Updates(models.Comment{Message: comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt,
	})

}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentID, _ := strconv.Atoi(c.Param("commentID"))

	var comment models.Comment

	err := db.Model(&comment).Delete(&comment, commentID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
