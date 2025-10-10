package memory

import (
	"context"
	"fmt"
	"main/internal/services/memory"
	use_case "main/internal/usecase/memory"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *use_case.UseCase
}

func NewController(useCase *use_case.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (uc Controller) CreateMemoryInCabinet(c *gin.Context) {
	var data memory.Create
	ctx := context.Background()
	authHeader := c.GetHeader("Authorization")

	data.FullName = c.PostForm("full_name")
	data.BirthPlace = c.PostForm("birth_place")
	data.BirthDate = c.PostForm("birth_date")
	data.BioHeadline = c.PostForm("bio_headline")
	data.Bio = c.PostForm("bio")
	data.FamilyMemberId, _ = strconv.Atoi(c.PostForm("family_member_id"))
	data.BioHeadline = c.PostForm("bio_headline")
	data.Bio = c.PostForm("bio")
	data.RegionId, _ = strconv.Atoi(c.PostForm("region_id"))
	data.DistrictId, _ = strconv.Atoi(c.PostForm("district_id"))
	data.CemeteryId, _ = strconv.Atoi(c.PostForm("cemetery_id"))
	data.DeathPlace = c.PostForm("death_place")
	data.DeathCauseId, _ = strconv.Atoi(c.PostForm("death_cause_id"))
	data.DeathDate = c.PostForm("death_date")
	data.MemoryStatus, _ = strconv.ParseBool(c.PostForm("memory_status"))

	file, _ := c.FormFile("avatar")
	if file != nil {
		filePath, err := uc.useCase.Upload(ctx, file, "./media/memory/avatar")
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}
		data.Avatar = filePath
	}

	form, err := c.MultipartForm()
	if err == nil {
		// Images
		if files := form.File["images"]; len(files) > 0 {
			paths, err := uc.useCase.MultipleUpload(ctx, files, "./media/memory/images")
			if err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}
			data.Images = paths
		}

		if files := form.File["videos"]; len(files) > 0 {
			paths, err := uc.useCase.MultipleUpload(ctx, files, "./media/memory/videos")
			if err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}
			data.Videos = paths
		}

		if files := form.File["audio"]; len(files) > 0 {
			paths, err := uc.useCase.MultipleUpload(ctx, files, "./media/memory/audio")
			if err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}
			data.Audio = paths
		}

		if vals := form.Value["social_id"]; len(vals) > 0 {
			for _, v := range vals {
				id, _ := strconv.Atoi(v)
				data.SocialId = append(data.SocialId, id)
			}
		}
	}

	detail, err := uc.useCase.CreateMemoryInCabinet(ctx, data, authHeader)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Memory created successfully",
		"data":    detail,
	})
}

func (uc Controller) GetListInCabinet(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	ctx := context.Background()
	list, err := uc.useCase.GetListInCabinet(ctx, authHeader)

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

func (uc Controller) DeleteInCabinet(c *gin.Context) {
	paramId := c.Param("id")
	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "param id mast be int",
		})
		return
	}

	ctx := context.Background()
	if err := uc.useCase.DeleteInCabinet(ctx, paramIdInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}

func (uc Controller) GetByIdInCabinet(c *gin.Context) {
	paramId := c.Param("id")
	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "param id mast be int",
		})
		return
	}

	ctx := context.Background()
	detail, err := uc.useCase.GetById(ctx, paramIdInt)

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

func (uc Controller) UpdateInCabinet(c *gin.Context) {
	var data memory.Update

	paramId := c.Param("id")
	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "param id must be int"})
		return
	}

	ctx := context.Background()

	fullName := c.PostForm("full_name")
	if fullName != "" {
		data.FullName = &fullName
	}

	birthPlace := c.PostForm("birth_place")
	if birthPlace != "" {
		data.BirthPlace = &birthPlace
	}

	birthDate := c.PostForm("birth_date")
	if birthDate != "" {
		data.BirthDate = &birthDate
	}

	bioHeadline := c.PostForm("bio_headline")
	if bioHeadline != "" {
		data.BioHeadline = &bioHeadline
	}

	bio := c.PostForm("bio")
	if bio != "" {
		data.Bio = &bio
	}

	deathPlace := c.PostForm("death_place")
	if deathPlace != "" {
		data.DeathPlace = &deathPlace
	}

	deathDate := c.PostForm("death_date")
	if deathDate != "" {
		data.DeathDate = &deathDate
	}

	if val := c.PostForm("memory_status"); val != "" {
		boolVal := val == "true" || val == "1"
		data.MemoryStatus = &boolVal
	}

	if val := c.PostForm("death_cause_id"); val != "" {
		id, _ := strconv.Atoi(val)
		data.DeathCauseId = &id
	}

	if val := c.PostForm("region_id"); val != "" {
		id, _ := strconv.Atoi(val)
		data.RegionId = &id
	}

	if val := c.PostForm("district_id"); val != "" {
		id, _ := strconv.Atoi(val)
		data.DistrictId = &id
	}

	if val := c.PostForm("cemetery_id"); val != "" {
		id, _ := strconv.Atoi(val)
		data.CemeteryId = &id
	}

	if val := c.PostForm("family_member_id"); val != "" {
		id, _ := strconv.Atoi(val)
		data.FamilyMemberId = &id
	}

	//delete img
	deleteImgs := c.PostFormArray("delete_img")
	if len(deleteImgs) > 0 {
		data.DeleteImg = &deleteImgs
	}
	if data.DeleteImg != nil {
		fmt.Println("➜⇒ ➜⇒ ➜⇒ data images", data.DeleteImg)
		for _, imgURL := range *data.DeleteImg {
			if err := uc.useCase.DeleteFile(ctx, imgURL); err != nil {
				fmt.Println("Failed to delete image:", imgURL, err)
			}
		}
	}

	//delete video
	deleteVideo := c.PostFormArray("delete_video")
	if len(deleteVideo) > 0 {
		data.DeleteVideo = &deleteVideo
	}
	if data.DeleteVideo != nil {
		fmt.Println("➜⇒ ➜⇒ ➜⇒ data images", data.DeleteVideo)
		for _, videoURL := range *data.DeleteVideo {
			if err := uc.useCase.DeleteFile(ctx, videoURL); err != nil {
				fmt.Println("Failed to delete image:", videoURL, err)
			}
		}
	}

	//delete audio
	deleteAudio := c.PostFormArray("delete_audio")
	if len(deleteAudio) > 0 {
		data.DeleteAudio = &deleteAudio
	}
	if data.DeleteAudio != nil {
		fmt.Println("➜⇒ ➜⇒ ➜⇒ data images", data.DeleteAudio)
		for _, audioURL := range *data.DeleteAudio {
			if err := uc.useCase.DeleteFile(ctx, audioURL); err != nil {
				fmt.Println("Failed to delete image:", audioURL, err)
			}
		}
	}

	file, _ := c.FormFile("avatar")

	if file != nil {
		filePath, err := uc.useCase.Upload(ctx, file, "./media/memory/avatar")
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}
		data.Avatar = &filePath
	}

	form, err := c.MultipartForm()

	//append img
	if err == nil && form != nil {
		if existing := form.Value["images_existing"]; len(existing) > 0 {
			data.Images = append(data.Images, existing...)
		}
	}

	//append video
	if err == nil && form != nil {
		if existing := form.Value["videos_existing"]; len(existing) > 0 {
			data.Videos = append(data.Videos, existing...)
		}
	}

	//append audio
	if err == nil && form != nil {
		if existing := form.Value["audio_existing"]; len(existing) > 0 {
			data.Audio = append(data.Audio, existing...)
		}
	}

	if err == nil {
		if files := form.File["images"]; len(files) > 0 {
			paths, err := uc.useCase.MultipleUpload(ctx, files, "./media/memory/images")
			if err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}
			data.Images = append(data.Images, paths...)
		}

		if files := form.File["videos"]; len(files) > 0 {
			paths, err := uc.useCase.MultipleUpload(ctx, files, "./media/memory/videos")
			if err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}
			data.Videos = append(data.Videos, paths...)
		}

		if files := form.File["audio"]; len(files) > 0 {
			paths, err := uc.useCase.MultipleUpload(ctx, files, "./media/memory/audio")
			if err != nil {
				c.JSON(400, gin.H{"message": err.Error()})
				return
			}
			data.Audio = append(data.Audio, paths...)
		}

		if vals := form.Value["social_id"]; len(vals) > 0 {
			for _, v := range vals {
				id, _ := strconv.Atoi(v)
				data.SocialId = append(data.SocialId, id)
			}
		}
	}
	detail, err := uc.useCase.UpdateInCabinet(ctx, data, paramIdInt)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok!",
		"data":    detail,
	})
}
