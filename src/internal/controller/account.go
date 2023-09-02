// https://documenter.getpostman.com/view/795261/2s9Xy6pUdQ#d6e264ba-8bdf-49e6-9a95-cd8c6b6d0126

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ntfargo/tir-goapi/src/internal/models"
	"github.com/ntfargo/tir-goapi/src/internal/others"
	"github.com/ntfargo/tir-goapi/src/internal/utils"
)

type AppError struct {
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func RegisterUser(c *gin.Context) {
	var user models.AuthCredentials
	if err := c.BindJSON(&user); err != nil {
		others.ErrorHandler(others.ErrInvalidInput, c.Writer, c.Request)
		return
	}

	var hashedPassword string

	isValid, errMsg := func() (bool, error) {
		isValid, errMsg := utils.ComplexValidator{}.Validate(user.Password)
		if !isValid {
			return false, AppError{Message: errMsg}
		}
		var err error
		hashedPassword, err = utils.BcryptHasher{}.Hash(user.Password)
		if err != nil {
			return false, err
		}
		return true, nil
	}()

	if !isValid {
		others.ErrorHandler(errMsg, c.Writer, c.Request)
		return
	}

	isValid, errMsg = utils.SimpleValidator{}.Validate(user.Email)
	if !isValid {
		others.ErrorHandler(errMsg, c.Writer, c.Request)
		return
	}

	userObject := models.User{
		Email:    user.Email,
		Password: hashedPassword,
	}

	newUserID, err := models.CreateUser(userObject)
	if err != nil {
		others.ErrorHandler(err, c.Writer, c.Request)
		return
	}

	c.JSON(200, gin.H{
		"message": "User registered successfully",
		"userID":  newUserID,
		"email":   user.Email,
	})
}
