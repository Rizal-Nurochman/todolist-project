package main

import (
	"time"
	"github.com/NurochmanRizal/go-ToDoList/database"
	"github.com/NurochmanRizal/go-ToDoList/handler"
	"github.com/NurochmanRizal/go-ToDoList/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	
	database.ConnectToDatabase()

	api:=router.Group("/api")
	{
		auth:=api.Group("/auth")
		{
			auth.POST("/signup", handler.SignUpHandler)
			auth.POST("/login", middleware.Auth, handler.LoginHandler)
		}
	}

	router.GET("/todo", handler.GetTodoByUser)
	router.GET("/todo:userid")
	router.POST("/todo", handler.PostToDoHandler)
}