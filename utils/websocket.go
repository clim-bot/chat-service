package utils

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "log"
    "net/http"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func HandleConnections(c *gin.Context) {
    ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    for {
        _, message, err := ws.ReadMessage()
        if err != nil {
            log.Println("Error reading message:", err)
            break
        }
        log.Printf("Received: %s", message)
    }
}
