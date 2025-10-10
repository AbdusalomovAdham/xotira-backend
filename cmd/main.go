package main

import (
	cache "main/internal/cache"
	auth_controller "main/internal/controllers/http/v1/auth"
	district_controller "main/internal/controllers/http/v1/districts"
	memory_controller "main/internal/controllers/http/v1/memory"
	region_controller "main/internal/controllers/http/v1/region"
	user_controller "main/internal/controllers/http/v1/user"
	auth_middleware "main/internal/middleware/auth"
	"main/internal/pkg/config"
	"main/internal/pkg/postgres"
	"main/internal/repository/postgres/districts"
	"main/internal/repository/postgres/memory"
	"main/internal/repository/postgres/region"
	"main/internal/repository/postgres/user"
	"main/internal/services/auth"
	districts_service "main/internal/services/districts"
	"main/internal/services/email"
	file_service "main/internal/services/file"
	memory_service "main/internal/services/memory"
	region_service "main/internal/services/region"
	user_service "main/internal/services/user"
	auth_use_case "main/internal/usecase/auth"
	district_use_case "main/internal/usecase/districts"
	memory_use_case "main/internal/usecase/memory"
	region_use_case "main/internal/usecase/region"
	user_use_case "main/internal/usecase/user"
	"time"
	video_service "main/internal/services/video"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	serverPost := ":" + config.GetConfig().Port

	r := gin.Default()

	//databases
	postgresDB := postgres.NewDB()

	r.Static("/media", "./media")

	//repositories
	userRepository := user.NewRepository(postgresDB)
	regionRepository := region.NewRepository(postgresDB)
	districtRepository := districts.NewRepository(postgresDB)
	memoryRepository := memory.NewRepository(postgresDB)

	//services
	authService := auth.NewService(userRepository)
	emailService := email.NewEmailService()
	userService := user_service.NewService(userRepository)
	fileService := file_service.NewService()
	videoService := video_service.NewService()
	regionService := region_service.NewService(regionRepository)
	districtService := districts_service.NewService(districtRepository)
	memoryService := memory_service.NewService(memoryRepository)

	//cache
	newCache := cache.NewCache(config.GetConfig().RedisHost, config.GetConfig().RedisDB, time.Duration(config.GetConfig().RedisExpires)*time.Second)

	//usecase
	authUseCase := auth_use_case.NewUseCase(authService, userRepository, newCache, emailService)
	userUseCase := user_use_case.NewUseCase(userService, authService, fileService)
	regionUseCase := region_use_case.NewUseCase(regionService)
	districtUseCase := district_use_case.NewUseCase(districtService)
	memoryUseCase := memory_use_case.NewUseCase(memoryService, authService, fileService,videoService)

	//controller
	authController := auth_controller.NewController(authUseCase)
	userController := user_controller.NewController(userUseCase)
	regionController := region_controller.NewController(regionUseCase)
	districtController := district_controller.NewController(districtUseCase)
	memoryController := memory_controller.NewController(memoryUseCase)

	//middleware
	authMiddleware := auth_middleware.NewMiddleware(authService)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081", "http://192.168.1.120:8081", "http://172.20.10.5:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		v1 := api.Group("v1")

		// #auth

		//sign-in
		v1.POST("/sign-in", authController.SignIn)
		//sign up
		v1.POST("/sign-up", authController.SignUp)
		//forgot password
		v1.POST("/forgot-psw", authController.ForgotPsw)
		//check code
		v1.POST("/check-code", authController.CheckCode)
		//update password
		v1.PATCH("/update-psw", authController.UpdatePsw)
		//resend code
		v1.POST("/resend-code", authController.ResendCode)

		//	#user
		//get-list
		v1.GET("/admin/user/list", authMiddleware.AuthMiddleware(), userController.AdminGetUserList)
		//get-detail
		v1.GET("admin/user/:id", authMiddleware.AuthMiddleware(), userController.AdminGetUserDetail)
		//update
		v1.PUT("admin/user/:id", authMiddleware.AuthMiddleware(), userController.AdminUpdateUser)
		//	create
		v1.POST("/amdin/user/create", authMiddleware.AuthMiddleware(), userController.AdminCreateUser)
		//delete
		v1.DELETE("admin/user/:id", authMiddleware.AuthMiddleware(), userController.AdminDeleteUser)
		//	get-by-email
		v1.GET("/get/user", authMiddleware.AuthMiddleware(), userController.GetByEmailWithLocation)

		//	#region
		//	get-list
		v1.GET("/region/list", regionController.GetAllRegion)

		//	 #district
		//	get-list
		v1.POST("/district/:id", districtController.GetAllRegion)

		//	#cabinet
		// update
		v1.PATCH("/update/cabinet", authMiddleware.AuthMiddleware(), userController.UpdateCabiner)

		// #memory-cabinet
		//get-list
		v1.GET("/list/memory", authMiddleware.AuthMiddleware(), memoryController.GetListInCabinet)
		//get-by-id
		v1.GET("/get/memory/:id", authMiddleware.AuthMiddleware(), memoryController.GetByIdInCabinet)
		//	create
		v1.POST("/create/memory", authMiddleware.AuthMiddleware(), memoryController.CreateMemoryInCabinet)
		//update
		v1.PATCH("update/memory/:id", authMiddleware.AuthMiddleware(), memoryController.UpdateInCabinet)
		//	delete
		v1.DELETE("/delete/memory/:id", authMiddleware.AuthMiddleware(), memoryController.DeleteInCabinet)
		// video-upload
		v1.POST("/upload/video", authMiddleware.AuthMiddleware(), memoryController.VideoUpload)

	}

	r.Run(serverPost)

}
