package controller

import "github.com/gin-gonic/gin"

func RootController() *gin.Engine {
	r := gin.Default()

	return r
}
