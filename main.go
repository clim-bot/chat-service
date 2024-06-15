package main

import (
    "fmt"
    "log"
    "os"

    "github.com/clim-bot/chat-service/routes"
    "github.com/clim-bot/chat-service/utils"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)


func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Print all relevant environment variables
    fmt.Println("AUTH0_DOMAIN:", os.Getenv("AUTH0_DOMAIN"))
    fmt.Println("AUTH0_CLIENT_ID:", os.Getenv("AUTH0_CLIENT_ID"))
    fmt.Println("AUTH0_CLIENT_SECRET:", os.Getenv("AUTH0_CLIENT_SECRET"))
    fmt.Println("AUTH0_AUDIENCE:", os.Getenv("AUTH0_AUDIENCE"))
    fmt.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))
    fmt.Println("SESSION_SECRET:", os.Getenv("SESSION_SECRET"))

	auth0Domain := os.Getenv("AUTH0_DOMAIN")
    if auth0Domain == "" {
        log.Fatal("AUTH0_DOMAIN is not set in the environment")
    }

    _, sqlDB := utils.SetupDatabase()
    defer sqlDB.Close()

    r := gin.Default()

    store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
    r.Use(sessions.Sessions("mysession", store))

    routes.SetupAuthRoutes(r)
    routes.SetupChatRoutes(r)

    r.Run(":8080")
}
