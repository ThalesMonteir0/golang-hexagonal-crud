package main

import (
	"github.com/ThalesMonteir0/golang-hexagonal-crud/adapter/input/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {

	}
	
	app := gin.Default()
	routes.InitRoutes(app)

}
