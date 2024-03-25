package router

import (
	"github.com/gin-gonic/gin"
	"project_akhir/app"
	"project_akhir/controller"
	"project_akhir/middleware"
	"project_akhir/repository"
	"project_akhir/service"
)

func PhotoRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewPhotoRepository(db)
	srv := service.NewPhotoService(repo)
	ctrl := controller.NewPhotoController(srv)

	photoRouter := router.Group("/photos", middleware.AuthMiddleware())

	{
		photoRouter.POST("/", ctrl.Create)
		photoRouter.GET("/", ctrl.GetAll)
		{
			photoRouter.PUT("/:id", middleware.PhotoMiddleware(srv), ctrl.Update)
			photoRouter.DELETE("/:id", middleware.PhotoMiddleware(srv), ctrl.Delete)
		}
	}
}
