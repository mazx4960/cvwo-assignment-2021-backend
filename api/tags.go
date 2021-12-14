package api

import (
	"errors"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTag(c *gin.Context) {
	id := c.Params.ByName("id")
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unable to parse jwt token"})
		return
	}

	var tag models.Tag
	err := database.DB.First(&tag, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "Record Not Found"})
		return
	}
	// Check if user is the owner of the tag
	if tag.UserID != userId {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	c.JSON(200, tag)
}

func GetTagsByUser(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unable to parse jwt token"})
		return
	}
	
	var tags []models.Tag
	err := database.DB.Where("user_id = ?", userId).Find(&tags).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "Record Not Found"})
		return
	}
	
	c.JSON(200, tags)
}

func CreateTag(c *gin.Context) {
	var tag models.Tag
	c.BindJSON(&tag)
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unable to parse jwt token"})
		return
	}

	tag.UserID = userId.(uint)
	database.DB.Create(&tag)
	c.JSON(http.StatusOK, gin.H{"tag": tag})
}

func DeleteTag(c *gin.Context) {
	id := c.Params.ByName("id")
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unable to parse jwt token"})
		return
	}

	var tag models.Tag
	err := database.DB.First(&tag, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "Record Not Found"})
		return
	}
	// Check if user is the owner of the tag
	if tag.UserID != userId {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	database.DB.Delete(&tag)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func AddTagToTask(c *gin.Context) {
	var taskTag models.TaskTag
	c.BindJSON(&taskTag)
	database.DB.Create(&taskTag)
	c.JSON(http.StatusOK, gin.H{"taskTag": taskTag})
}

func RemoveTagFromTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var taskTag models.TaskTag
	err := database.DB.First(&taskTag, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "Record Not Found"})
		return
	}

	database.DB.Delete(&taskTag)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func GetTagsByTask(c *gin.Context) {
	id := c.Params.ByName("task_id")
	var taskTags []models.TaskTag
	database.DB.Where("task_id = ?", id).Find(&taskTags)

	var tags []models.Tag
	for _, taskTag := range taskTags {
		var tag models.Tag
		database.DB.First(&tag, taskTag.TagID)
		tags = append(tags, tag)
	}
	c.JSON(200, tags)
}

func GetTasksByTag(c *gin.Context) {
	id := c.Params.ByName("tag_id")
	var taskTags []models.TaskTag
	database.DB.Where("tag_id = ?", id).Find(&taskTags)

	var tasks []models.Task
	for _, taskTag := range taskTags {
		var task models.Task
		database.DB.First(&task, taskTag.TaskID)
		tasks = append(tasks, task)
	}
	c.JSON(200, tasks)
}