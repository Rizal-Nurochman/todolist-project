package handler

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/NurochmanRizal/go-ToDoList/database"
	"github.com/NurochmanRizal/go-ToDoList/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(c *gin.Context)  {
	var userBody struct{
		Email string
		Password string		
	}

	err:=c.ShouldBindJSON(&userBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	hash, err:=bcrypt.GenerateFromPassword([]byte(userBody.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	NewUser:=models.User{
		Email: userBody.Email,
		Password: string(hash),
	}
	err = database.DB.Create(&NewUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"Create account success!",
	})
}

func LoginHandler(c *gin.Context)  {
	var userBody struct{
		Email string
		Password string
	}

	err:=c.ShouldBindJSON(&userBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	var User models.User
	err=database.DB.Where("email=?", userBody.Email).First(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid email or password",
		})
		return
	}

	err=bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(userBody.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": User.ID,
	"exp": time.Now().Add(time.Hour*24*1).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("SECRET"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message":"Login success!",
		"data":"welcome "+User.Email,
	})
}

func Validator(c *gin.Context) {
    user,_ := c.Get("user")
		c.JSON(http.StatusOK, gin.H{
			"user":user,
		})
}

func PostToDoHandler(c *gin.Context) {
	var todos models.Todos
	err:=c.ShouldBindJSON(&todos)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":err.Error(),
		})
		return
	}

	userID:=c.MustGet("userid").(uint)
	newTodo:= models.Todos{
		Name: todos.Name,
		UserID: userID,
		IsDone: todos.IsDone,
	}

	err=database.DB.Create(&newTodo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messages":"Create to do list Success!",
		"data":newTodo,
	})
}

func GetTodoByUser(c *gin.Context)  {
	userID:=c.MustGet("userid").(uint)
	var todo []models.Todos
	err:=database.DB.Where("userid = ?", userID).Find(&todo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":todo,
	})
}

func GetTodoByIDUser(c *gin.Context) {
	idString:=c.Param("id")
	id, err:=strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	var todo models.Todos
	userID:=c.MustGet("userid").(uint)
	err=database.DB.Where("id = ? AND userid = ?", id, userID).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":todo,
	})

}