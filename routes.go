package main

import (
	"main/api"
	"main/graph/resolver"
	"main/graph/generated"
	"main/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func initializeRoutesv1(r *gin.Engine) {
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

func initializeRoutesv2(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())

	public := r.Group("/apiv2")
	protected := r.Group("/apiv2")

	// Middleware
	protected.Use(middleware.AuthMiddleware())
	protected.Use(middleware.GinContextToContextMiddleware())

	// User routes
  public.POST("/login", api.Login)
  public.POST("/register", api.Register)
  protected.POST("/logout", api.Logout)
	protected.GET("/user", api.GetUser)

	protected.POST("/query", graphqlHandler())
	public.GET("/", playgroundHandler())
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}