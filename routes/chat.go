package routes

import (
    "github.com/clim-bot/chat-service/controllers"
    "github.com/gin-gonic/gin"
)

func SetupChatRoutes(r *gin.Engine) {
    chat := r.Group("/chat")
    {
        chat.GET("/ws", controllers.ChatWebSocket)
    }
}
