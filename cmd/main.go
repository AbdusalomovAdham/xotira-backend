package main

import (
	"fmt"
	cache "main/internal/cache"
	auth_controller "main/internal/controllers/http/v1/auth"
	"main/internal/pkg/config"
	"main/internal/pkg/postgres"
	"main/internal/repository/postgres/user"
	"main/internal/services/auth"
	"main/internal/services/email"
	auth_use_case "main/internal/usecase/auth"
	"time"

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

	//cache
	newCache := cache.NewCache(config.GetConfig().RedisHost, config.GetConfig().RedisDB, time.Duration(config.GetConfig().RedisExpires)*time.Second)

	//usecase
	authUseCase := auth_use_case.NewUseCase(authService, userRepository, newCache, emailService)

	//controller
	userController := auth_controller.NewController(authUseCase)

	//middleware
	//authMiddleware := auth_middleware.NewMiddleware(authService)

	//auth
	r.POST("/sign-in", userController.SignIn)
	r.POST("/sign-up", userController.SignUp)
	r.POST("/forgot-psw", userController.ForgotPsw)
	r.POST("/check-code", userController.CheckCode)
	r.PATCH("/update-psw", userController.UpdatePsw)
	r.Run(serverPost)
}
