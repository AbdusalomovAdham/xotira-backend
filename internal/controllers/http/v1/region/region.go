package region

import (
	"context"
	region_service "main/internal/services/region"
	"main/internal/usecase/region"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *region.UseCase
}

func NewController(useCase *region.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (uc Controller) GetAllRegion(c *gin.Context) {
	var filter region_service.Filter
	query := c.Request.URL.Query()
	lang := c.GetHeader("Accept-Language")
	limitQ := query["limit"]
	if len(limitQ) > 0 {
		queryInt, err := strconv.Atoi(limitQ[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Limit must be number!",
			})

			return
		}

		filter.Limit = &queryInt
	}

	offsetQ := query["offset"]
	if len(offsetQ) > 0 {
		queryInt, err := strconv.Atoi(offsetQ[0])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Offset must be number!",
			})

			return
		}
		filter.Offset = &queryInt
	}

	ctx := context.Background()

	orderParam := c.Query("order")
	if orderParam == "" {
		orderParam = "id asc"
	}

	parts := strings.Fields(orderParam)

	direction := "asc"
	if len(parts) > 1 {
		direction = parts[1]
	}

	list, count, err := uc.useCase.GetAll(ctx, filter, direction, lang)
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
			"count":   count,
		},
	})

}
