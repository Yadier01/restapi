package routes

import (
	"net/http"

	"github.com/Yadier01/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		key := c.GetHeader("Auth")
		s, err := utils.JwtParse(key)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "could not parse token"})
			return
		}
		c.Set("id", s.Id)
		c.Next()
	}

}
