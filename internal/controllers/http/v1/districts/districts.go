package districts

import (
	"context"
	"main/internal/usecase/districts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *districts.UseCase
}

func NewController(useCase *districts.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (uc Controller) GetAllRegion(c *gin.Context) {

	regionIDStr := c.Param("id")
	region_id, err := strconv.Atoi(regionIDStr)

	lang := c.GetHeader("Accept-Language")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid region_id",
		})
		return
	}

	ctx := context.Background()
	list, err := uc.useCase.GetByRegionId(ctx, region_id, lang)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"data": map[string]interface{}{
			"results": list,
		},
	})

}
