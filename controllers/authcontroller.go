package controllers

import (
	"api-starter/models/entities"
	"api-starter/models/service"
	"fmt"
	"log"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	userService service.UserService
}

func (auth *AuthController) Login(c *gin.Context) {

	var loginInfo entities.User
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, errf := auth.userService.Find(&loginInfo)
	if errf != nil {
		c.JSON(401, gin.H{"error": "Not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		c.JSON(402, gin.H{"error": "Email or password is invalid."})
		return
	}

	fmt.Println("user email is ", user.Email)
	token, err := user.GetJwtToken()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func (auth *AuthController) GetUser(c *gin.Context) {
	temporaryUser := &entities.User{Email: c.Param("email")}
	user, errf := auth.userService.Find(temporaryUser)
	if errf != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	s := structs.New(user)
	s.TagName = "json"
	userMap := s.Map()
	delete(userMap, "password")
	c.JSON(200, userMap)
}

func (auth *AuthController) UpdateUser(c *gin.Context) {
	var tempUser *entities.User
	if err := c.ShouldBindJSON(&tempUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := auth.userService.Update(tempUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": "ok"})
}

func (auth *AuthController) Signup(c *gin.Context) {

	type signupInfo struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name"`
	}
	var info signupInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(401, gin.H{"error": "Please input all fields"})
		return
	}
	user := entities.User{}
	user.Email = info.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.Password = string(hash)
	user.Name = info.Name
	if err = auth.userService.Create(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": "ok"})
	}
	return
}
