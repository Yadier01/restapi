package routes

import (
	"net/http"

	"github.com/Yadier01/rest-api/models"
	"github.com/Yadier01/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not bind json"})
		return
	}
	err := user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}
func userLogin(c *gin.Context) {
	var users models.Users
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not bind json"})
		return
	}
	u, err := models.GetUserByEmail(users.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	err = utils.ComparePassword(u.Password, users.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
	}

	token, err := utils.JwtNew(u.ID, u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token", "err": err.Error()})
		return
	}
	c.Header("Auth", token)
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}
