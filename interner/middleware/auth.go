package middleware

import (
	"context"
	"net/http"

	"github.com/Jangwooo/2022Hackathon/interner/domain/object/response"
	"github.com/Jangwooo/2022Hackathon/interner/redis"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")

	r := redis.Connection()

	userID, err := r.Get(context.Background(), token).Result()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, response.Error{Massage: "Token is not available"})
		return
	}

	c.Set("userID", userID)
	c.Next()
}
