package auth

import (
	"context"
	auth_service "main/internal/services/auth"
	"main/internal/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *auth.UseCase
}

func NewController(useCase *auth.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (ac Controller) SignIn(c *gin.Context) {
	var data auth_service.SignIn
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx := context.Background()
	token, err := ac.useCase.SignIn(ctx, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (as Controller) SignUp(c *gin.Context) {
	var data auth_service.SignUp
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx := context.Background()
	token, err := as.useCase.SignUp(ctx, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (ac Controller) ForgotPsw(c *gin.Context) {
	var data auth_service.ForgotPsw
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx := context.Background()
	token, err := ac.useCase.ForgotPsw(ctx, data.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (as Controller) CheckCode(c *gin.Context) {
	var data auth_service.CheckCode
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx := context.Background()
	err = as.useCase.CheckCode(ctx, data.Code, data.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "the code is correct"})
}

func (as Controller) UpdatePsw(c *gin.Context) {
	var data auth_service.UpdatePsw

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx := context.Background()
	err = as.useCase.ResetPsw(ctx, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password update"})
}
