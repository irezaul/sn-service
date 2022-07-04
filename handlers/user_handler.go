package handlers

import (
	"ginorm/database"
	"ginorm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"Api":  "Test_api",
		"name": "rezaul",
	})

}

func AddUser(c *gin.Context) {
	var User model.User
	DB := database.DB
	err := c.BindJSON(&User)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	DB.Create(&User)
	c.JSON(http.StatusOK, User)
}

func GetUser(c *gin.Context) {
	var User []model.User
	DB := database.DB
	DB.Find(&User)
	c.JSON(http.StatusOK, &User)
}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var User model.User
	DB := database.DB
	err := DB.Where("id=?", id).First(&User).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	DB.Delete(&User)
	c.JSON(http.StatusOK, gin.H{
		"Status": "Delete Successfully",
	})
}

func UpdateUser(c *gin.Context) {
	var User model.User
	id := c.Param("id")
	DB := database.DB
	err := DB.Where("id=?", id).Find(&User).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&User)
	DB.Save(&User)
	c.JSON(http.StatusOK, gin.H{
		"Status": "Update done....",
	})

}
