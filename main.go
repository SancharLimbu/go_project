package main

import (
	"log"
	"os"
	"time"

	"go-api/controllers"
	"go-api/database"
	"go-api/middleware"
	"go-api/routes"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	// TODO:
	// - [x] 1. Add EditProductAdmin
	// 2. Add OTP with https://github.com/AfterShip/email-verifier.git
	// 3. Add Admin authentication
	// - [x] 4. Add password and email regex

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()

	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type, Token",
		ExposedHeaders: "",
		MaxAge:         50 * time.Second,
		// Dont use Origins: "*" when using Credentials: true
		Credentials:     false,
		ValidateHeaders: false,
	}))

	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
