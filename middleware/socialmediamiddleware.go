package middleware

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
	"project_akhir/model/domain"
	"project_akhir/model/http/response"
	"project_akhir/service"
)

func SocialMediaMiddleware(socialMediaService service.SocialMediaService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			socialMedia domain.SocialMedia
			err         error
		)

		socialMediaId, _ := strconv.Atoi(ctx.Param("id"))
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		if socialMedia, err = socialMediaService.GetOne(uint(socialMediaId)); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Errors: gin.H{
					"message": "Social media not found",
				},
			})

			return
		}

		if socialMedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
				Code:   http.StatusForbidden,
				Status: "Forbidden",
				Errors: gin.H{
					"message": "You don't have permission",
				},
			})

			return
		}
	}
}