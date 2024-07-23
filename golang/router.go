package main

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	router.POST("/api/strong_password_steps", getPasswordStrength)
}
