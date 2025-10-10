package gallery

import (
	"context"
	"main/internal/usecase/gallery"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *gallery.UseCase
}

func NewController(useCase *gallery.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (uc Controller) Create(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image not provided"})
		return
	}

	ctx := context.Background()
	imgPath, err := uc.useCase.Upload(ctx, file, "./media/gallery")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := uc.useCase.Create(ctx, imgPath, authHeader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
	})
}

func (uc Controller) GetAll(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	ctx := context.Background()
	list, err := uc.useCase.GetAll(ctx, authHeader)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"data":    list,
	})
}

type Delete struct {
	Image string `json:"image"`
}

func (uc Controller) DeleteInCabinet(c *gin.Context) {
	var url Delete
	paramStr := c.Param("id")
	id, err := strconv.Atoi(paramStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.ShouldBind(&url); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	ctx := context.Background()
	if err := uc.useCase.Delete(ctx, id, url.Image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok!"})

}
