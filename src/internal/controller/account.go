package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ntfargo/tir-goapi/src/internal/models"
	"github.com/ntfargo/tir-goapi/src/internal/others"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		w := c.Writer
		others.ErrorHandler(err, w)
		return
	}
}
