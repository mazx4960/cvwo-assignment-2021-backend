package main

import (
	"main/api"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())

	public := r.Group("/api")
	protected := r.Group("/api")

	// Middleware
	protected.Use(middleware.AuthMiddleware())

	// for testing
	public.GET("/", index)

	// User routes
  public.POST("/login", api.Login)
  public.POST("/register", api.Register)
  protected.POST("/logout", api.Logout)
	protected.GET("/user", api.GetUser)

	// Task routes
	protected.GET("/tasks", api.GetTaskByUser)
	protected.GET("/tasks/:id", api.GetTask)
	protected.POST("/tasks", api.CreateTask)
	protected.PUT("/tasks/:id", api.UpdateTask)
	protected.DELETE("/tasks/:id", api.DeleteTask)

	// Tags routes
	protected.GET("/tags", api.GetTagsByUser)
	protected.GET("/tags/:id", api.GetTag)
	protected.POST("/tags", api.CreateTag)
	protected.DELETE("/tags/:id", api.DeleteTag)

	// TaskTags routes
	protected.GET("/tasktags/tags/:task_id", api.GetTagsByTask)
	protected.GET("/tasktags/tasks/:tag_id", api.GetTasksByTag)
	protected.POST("/tasktags", api.AddTagToTask)
	protected.DELETE("/tasktags/:id", api.RemoveTagFromTask)
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the CVWO Assignment 2020 Backend",
	})
}