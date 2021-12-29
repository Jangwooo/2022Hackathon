package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func SetUp() *gin.Engine {
	r := gin.Default()

	return r
}
