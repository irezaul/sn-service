package handlers

import (
	"ginorm/database"
	"ginorm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddService(c *gin.Context) {
	var Service model.Service
	DB := database.DB
	err := c.BindJSON(&Service)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	DB.Create(&Service)
	c.JSON(http.StatusOK, Service)
}

func GetService(c *gin.Context) {
	var Service []model.Service
	DB := database.DB
	DB.Find(&Service)
	c.JSON(http.StatusOK, &Service)

}

func GetServiceById(c *gin.Context) {
	var Service model.Service
	id := c.Param("id")
	DB := database.DB
	err := DB.Where("id=?", id).Find(&Service).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, Service)
}

func DeleteService(c *gin.Context) {
	id := c.Param("id")
	var Service model.Service
	DB := database.DB
	err := DB.Where("id=?", id).First(&Service).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	DB.Delete(&Service)
	c.JSON(http.StatusOK, gin.H{
		"Status": "Delete Successfully",
	})
}
func UpdateService(c *gin.Context) {
	var Service model.Service
	id := c.Param("id")
	DB := database.DB
	err := DB.Where("id=?", id).Find(&Service).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&Service)
	DB.Save(&Service)
	c.JSON(http.StatusOK, gin.H{
		"Status": "Update done....",
	})
}
