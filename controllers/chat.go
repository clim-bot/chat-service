package controllers

import (
    "github.com/clim-bot/chat-service/utils"
    "github.com/gin-gonic/gin"
)

func ChatWebSocket(c *gin.Context) {
    utils.HandleConnections(c)
}
