package controllers_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/clim-bot/chat-service/controllers"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
    router := gin.Default()
    router.GET("/auth/login", controllers.Login)

    req, _ := http.NewRequest("GET", "/auth/login", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
}
