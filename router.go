package main

import (
	"crowdfunding-api/src/modules/campaign"
	"crowdfunding-api/src/modules/category"
	"crowdfunding-api/src/modules/general/file"
	"crowdfunding-api/src/modules/user"
	"net/http"

	"github.com/gin-gonic/gin"

	"crowdfunding-api/src/middlewares"
	"crowdfunding-api/src/modules/auth"
)

// SetupRoutes Setup Routes
func SetupRoutes(app *gin.Engine) {
	// Check Server Status Endpoint
	app.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Server alive!",
			"data":    context,
		})
	})

	v1 := app.Group("/api/v1")
	{
		public := v1.Group("")
		{
			auth.NewRouter(public.Group("/auth"))
			category.NewPublicRouter(public.Group("/categories"))
			campaign.NewPublicRouter(public.Group("/campaigns"))
		}

		private := v1.Group("")
		{
			private.Use(middlewares.JwtAuthMiddleware())

			category.NewRouter(private.Group("/categories"))
			file.NewFileRouter(private.Group("/file"))
			campaign.NewRouter(private.Group("/campaigns"))
			user.NewRouter(private.Group("/users"))
		}
	}
}
