package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PasswordRequest struct {
	InitPassword string `json:"init_password"`
}

type PasswordResponse struct {
	NumOfSteps int `json:"num_of_steps"`
}

// Remove this if it is already declared in utils.go
// func calculateSteps(password string) int {
//     ...
// }

func getPasswordStrength(c *gin.Context) {
	var request PasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	steps := calculateSteps(request.InitPassword)
	response := PasswordResponse{NumOfSteps: steps}
	c.JSON(http.StatusOK, response)
}
