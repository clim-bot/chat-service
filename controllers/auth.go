package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/clim-bot/chat-service/utils"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var (
	clientID     string
	clientSecret string
	redirectURL  = "http://localhost:8080/auth/callback"
	issuer       string

	provider *oidc.Provider
	verifier *oidc.IDTokenVerifier
	config   *oauth2.Config
)

type Auth0User struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Connection string `json:"connection"`
}

func init() {
	clientID = os.Getenv("AUTH0_CLIENT_ID")
	clientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	issuer = "dev-s1z0qdvc7odc3jc5.us.auth0.com"
	// issuer       = os.Getenv("AUTH0_DOMAIN")

	var err error
	provider, err = oidc.NewProvider(context.Background(), "https://"+issuer+"/")
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}

	config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	verifier = provider.Verifier(&oidc.Config{ClientID: clientID})
}

func Login(c *gin.Context) {
	state := "random" // In a real app, use a random generated state
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, config.AuthCodeURL(state))
}

func Callback(c *gin.Context) {
	session := sessions.Default(c)
	if c.Query("state") != session.Get("state") {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := config.Exchange(context.Background(), c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	idToken, err := verifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var claims struct {
		Email string `json:"email"`
	}
	if err := idToken.Claims(&claims); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session.Set("user", claims.Email)
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func Register(c *gin.Context) {
	var newUser Auth0User
	if err := c.BindJSON(&newUser); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	log.Printf("Received registration request: %+v", newUser)

	newUser.Connection = "Username-Password-Authentication"

	managementToken, err := utils.GetAuth0ManagementToken()
	if err != nil {
		log.Printf("Error getting Auth0 management token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get management token"})
		return
	}

	auth0Domain := "dev-s1z0qdvc7odc3jc5.us.auth0.com"
	auth0APIURL := "https://" + auth0Domain + "/api/v2/users"

	userJSON, err := json.Marshal(newUser)
	if err != nil {
		log.Printf("Error marshalling user JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("POST", auth0APIURL, bytes.NewBuffer(userJSON))
	if err != nil {
		log.Printf("Error creating new request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+managementToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request to Auth0: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp.StatusCode != http.StatusCreated {
		log.Printf("Failed to create user: %s", resp.Status)
		log.Printf("Response body: %s", responseBody)
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to create user", "details": string(responseBody)})
		return
	}

	log.Println("User registered successfully")
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
