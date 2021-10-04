package main

import (
	"fmt"
	"log"

	"github.com/helipoc/goapi/database"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"github.com/helipoc/goapi/api/postsapi"
	"github.com/helipoc/goapi/api/usersapi"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Couldn't Load Env Variables !")
	}
	app := gin.Default()
	err := database.Connect()
	if err != nil {
		fmt.Println("could'nt connect to database")
		return
	}

	usersRoute := app.Group("/users")
	postsRoute := app.Group("/posts")
	usersapi.MountRoutes(usersRoute)
	postsapi.MountRoutes(postsRoute)

	app.Run(":8080")
}
