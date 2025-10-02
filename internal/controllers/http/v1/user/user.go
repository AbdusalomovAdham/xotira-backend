package user

import (
	"context"
	"fmt"
	user_service "main/internal/services/user"
	"main/internal/usecase/user"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *user.UseCase
}

func NewController(useCase *user.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (uc Controller) AdminGetUserList(c *gin.Context) {
	var filter user_service.Filter
	query := c.Request.URL.Query()

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

	field := parts[0]
	direction := "asc"
	if len(parts) > 1 {
		direction = parts[1]
	}

	fmt.Println("field:", field, "direction:", direction)

	list, count, err := uc.useCase.AdminGetUserList(ctx, filter, direction)
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

func (uc Controller) AdminCreateUser(c *gin.Context) {
	var data user_service.Create

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"message": "Authorization header required"})
		return
	}

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if file != nil {
		filePath, err := uc.useCase.Upload(ctx, file, "./media/user/avatar")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		fmt.Println("file", filePath)
		data.Avatar = filePath
		fmt.Println("data", data)
	}

	detail, err := uc.useCase.AdminCreateUser(ctx, data, authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok!", "data": detail})
}

func (uc Controller) AdminGetUserDetail(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Id must be number!",
		})
		return
	}

	ctx := context.Background()

	detail, err := uc.useCase.AdminGetUserDetail(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"data":    detail,
	})
}

func (uc Controller) AdminUpdateUser(c *gin.Context) {
	var data user_service.Update
	authHeader := c.GetHeader("Authorization")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Id must be a number"})
		return
	}
	data.Id = &id
	if data.FullName != nil {
		fmt.Println("data form", *data.FullName)
	} else {
		fmt.Println("data form is nil")
	}
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()

	file, _ := c.FormFile("avatar")

	if file != nil {
		filePath, err := uc.useCase.Upload(ctx, file, "./media/user/avatar")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		data.Avatar = &filePath
	}

	detail, err := uc.useCase.AdminUpdateUser(ctx, data, file != nil, authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": detail})
}

func (uc Controller) AdminDeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Id must be a number"})
		return
	}

	ctx := context.Background()

	if err := uc.useCase.AdminDeleteUser(ctx, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (uc Controller) GetByEmail(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	ctx := context.Background()
	detail, err := uc.useCase.GetByEmail(ctx, authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": detail})
}
