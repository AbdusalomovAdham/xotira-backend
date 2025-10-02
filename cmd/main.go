package main

import (
	"fmt"
	cache "main/internal/cache"
	auth_controller "main/internal/controllers/http/v1/auth"
	user_controller "main/internal/controllers/http/v1/user"
	auth_middleware "main/internal/middleware/auth"
	"main/internal/pkg/config"
	"main/internal/pkg/postgres"
	"main/internal/repository/postgres/user"
	"main/internal/services/auth"
	"main/internal/services/email"
	file_service "main/internal/services/file"
	user_service "main/internal/services/user"
	auth_use_case "main/internal/usecase/auth"
	user_use_case "main/internal/usecase/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	serverPost := ":" + config.GetConfig().Port

	r := gin.Default()
	fmt.Println("Connect post", serverPost)

	//databases
	postgresDB := postgres.NewDB()

	//repositories
	userRepository := user.NewRepository(postgresDB)

	//services
	authService := auth.NewService(userRepository)
	emailService := email.NewEmailSeervice()
	userService := user_service.NewService(userRepository)
	fileService := file_service.NewService()

	//cache
	newCache := cache.NewCache(config.GetConfig().RedisHost, config.GetConfig().RedisDB, time.Duration(config.GetConfig().RedisExpires)*time.Second)

	//usecase
	authUseCase := auth_use_case.NewUseCase(authService, userRepository, newCache, emailService)
	userUserCase := user_use_case.NewUseCase(userService, authService, fileService)

	//controller
	authController := auth_controller.NewController(authUseCase)
	userController := user_controller.NewController(userUserCase)

	//middleware
	authMiddleware := auth_middleware.NewMiddleware(authService)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081", "http://127.0.0.1:8081"}, // frontend URL(lar)ni yozing
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
		v1.GET("/get/user", authMiddleware.AuthMiddleware(), userController.GetByEmail)
	}

	r.Run(serverPost)

}
