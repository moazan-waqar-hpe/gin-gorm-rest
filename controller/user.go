package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/moazan-waqar-hpe/gin-gorm-rest/config"
	"github.com/moazan-waqar-hpe/gin-gorm-rest/model"
)

func GetUsers(c *gin.Context) {
	users := []model.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func GetUser(c *gin.Context) {
	var user model.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.JSON(200, &user)

}

func CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)

}

func DeleteUser(c *gin.Context) {
	var user model.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)

}

func UpdateUser(c *gin.Context) {
	var user model.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)

}
