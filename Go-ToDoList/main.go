package main

import (
	"github.com/NurochmanRizal/go-ToDoList/database"
	"github.com/NurochmanRizal/go-ToDoList/handler"
	"github.com/NurochmanRizal/go-ToDoList/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectToDatabase()

	api:=router.Group("/api")
	{
		auth:=api.Group("/auth")
		{
			auth.POST("/signup", handler.SignUpHandler)
			auth.POST("/login", middleware.Auth, handler.LoginHandler)
		}
	}

	router.

}