package router

import (
	"github.com/gin-gonic/gin"
	"project_akhir/app"
	"project_akhir/controller"
	"project_akhir/middleware"
	"project_akhir/repository"
	"project_akhir/service"
)

func SocialMediaRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewSocialMediaRepository(db)
	srv := service.NewSocialMediaService(repo)
	ctrl := controller.NewSocialMediaController(srv)

	socialMedia := router.Group("/socialmedias", middleware.AuthMiddleware())

	{
		socialMedia.GET("/", ctrl.GetAll)
		socialMedia.POST("/", ctrl.Create)
		{
			socialMedia.PUT("/:id", middleware.SocialMediaMiddleware(srv), ctrl.Update)
			socialMedia.DELETE("/:id", middleware.SocialMediaMiddleware(srv), ctrl.Delete)
		}
	}
}
