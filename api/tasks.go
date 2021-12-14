package api

import (
	"errors"
	"main/database"
	"main/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTask(c *gin.Context) {
	id := c.Params.ByName("id")
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unable to parse jwt token"})
		return
	}

	var task models.Task
	err := database.DB.First(&task, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "Record Not Found"})
		return
	}
	// Cannot get other user's task
	if task.UserID != userId.(uint) {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	c.JSON(200, task)
}

func GetTaskByUser(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unable to parse jwt token"})
		return
	}

	var tasks []models.Task
	err := database.DB.Where("user_id = ?", userId.(uint)).Find(&tasks).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "Record Not Found"})
	} else {
		c.JSON(200, tasks)
	}
}

func CreateTask(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unable to parse jwt token"})
		return
	}

	task.UserID = userId.(uint)
	task.Status = false
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	database.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Params.ByName("id")
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unable to parse jwt token"})
		return
	}

	var task models.Task
	err := database.DB.First(&task, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "Record Not Found"})
		return
	}
	// Cannot update other user's task
	if task.UserID != userId.(uint) {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	c.BindJSON(&task)
	task.UpdatedAt = time.Now()
	
	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Params.ByName("id")
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unable to parse jwt token"})
		return
	}

	var task models.Task
	err := database.DB.First(&task, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "Record Not Found"})
		return
	}
	// Cannot delete other user's task
	if task.UserID != userId.(uint) {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	database.DB.Delete(&task)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}