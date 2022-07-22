package main

import (
	"crowdfunding-api/src/cores/database"
	"crowdfunding-api/src/middlewares"
	"crowdfunding-api/src/modules/auth"
	"crowdfunding-api/src/modules/campaign"
	"crowdfunding-api/src/modules/category"
	"crowdfunding-api/src/modules/general/file"
	"crowdfunding-api/src/modules/user"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// Setup Database Connection
	db := database.SetupDatabase()
	database.MigrateDb(db)

	// initialize modules/apps
	auth.New(db).InitModule()
	category.New(db).InitModule()
	user.New(db).InitModule()
	campaign.New(db).InitModule()
	file.New().InitModule()

	// Get Gin Mode from ENV
	mode := os.Getenv("GIN_MODE")

	// Set Gin Mode
	gin.SetMode(mode)

	// Create New App Instance
	app := gin.Default()

	// Setup CORS
	// app.Use(cors.Default())
	app.Use(middlewares.CORSMiddleware())
	//app.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"https://foo.com"},
	//	AllowMethods:     []string{"PUT", "POST", "GET"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "https://github.com"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))

	// Setup Routes
	SetupRoutes(app)

	// Run App at 3000
	err := app.Run(":3000")
	if err != nil {
		return
	}
}
