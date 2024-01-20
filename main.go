package main

import (
	"fmt"

	"github.com/khuongnguyenBlue/slacky/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	e := echo.New()
	e.Use(middlewares.ErrorHandler)
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	app := initializeApp()
	auth := e.Group("/auth")
	auth.POST("/register", app.authController.Register)
	auth.POST("/login", app.authController.Login)
	port := viper.GetString("PORT")
	e.Logger.Fatal(e.Start(port))
}
