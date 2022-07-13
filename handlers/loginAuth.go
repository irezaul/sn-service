package handlers

import (
	"sn-service/database"
	"net/http"
	"sn-service/model"

	"github.com/gin-gonic/gin"
)

// var store = sessions.NewCookieStore([]byte("secret"))

func LoginAuth(c *gin.Context) {
	var login model.Login
	var user []model.User
	db := database.DB
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err1 := db.Where("email = ?", login.Email).Where("password= ?", login.Password).First(&user).Error

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "Login Sucess",
	})
}
