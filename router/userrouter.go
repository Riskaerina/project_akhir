package router

import (
	"github.com/gin-gonic/gin"
	"project_akhir/app"
	"project_akhir/controller"
	"project_akhir/middleware"
	"project_akhir/repository"
	"project_akhir/service"
)

func UserRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo)
	ctrl := controller.NewUserController(srv)

	userRouter := router.Group("/users")

	{
		userRouter.POST("/register", ctrl.Register)
		userRouter.POST("/login", ctrl.Login)

		{
			userRouter.PUT("/", middleware.AuthMiddleware(), ctrl.Update)
			userRouter.DELETE("/", middleware.AuthMiddleware(), ctrl.Delete)
		}
	}
}