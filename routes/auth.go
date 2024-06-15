package routes

import (
    "github.com/clim-bot/chat-service/controllers"
    "github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine) {
    auth := r.Group("/auth")
    {
        auth.GET("/login", controllers.Login)
        auth.GET("/callback", controllers.Callback)
        auth.GET("/logout", controllers.Logout)
		auth.POST("/register", controllers.Register) 
    }
}
