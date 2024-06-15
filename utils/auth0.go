package utils

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
)

func GetAuth0ManagementToken() (string, error) {
    auth0Domain := os.Getenv("AUTH0_DOMAIN")
    clientID := os.Getenv("AUTH0_CLIENT_ID")
    clientSecret := os.Getenv("AUTH0_CLIENT_SECRET")
    audience := os.Getenv("AUTH0_AUDIENCE")

    if auth0Domain == "" || clientID == "" || clientSecret == "" || audience == "" {
        log.Printf("Environment variables missing. AUTH0_DOMAIN: %s, AUTH0_CLIENT_ID: %s, AUTH0_CLIENT_SECRET: %s, AUTH0_AUDIENCE: %s",
            auth0Domain, clientID, clientSecret, audience)
        return "", fmt.Errorf("missing environment variables")
    }

    auth0APIURL := fmt.Sprintf("https://%s/oauth/token", auth0Domain)
    log.Printf("Auth0 API URL: %s", auth0APIURL)

    requestBody := map[string]string{
        "client_id":     clientID,
        "client_secret": clientSecret,
        "audience":      audience,
        "grant_type":    "client_credentials",
    }
    jsonValue, err := json.Marshal(requestBody)
    if err != nil {
        log.Printf("Error marshalling request body: %v", err)
        return "", err
    }

    req, err := http.NewRequest("POST", auth0APIURL, bytes.NewBuffer(jsonValue))
    if err != nil {
        log.Printf("Error creating new request: %v", err)
        return "", err
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Error making request to Auth0: %v", err)
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        bodyBytes, _ := io.ReadAll(resp.Body)
        log.Printf("Failed to get management token, status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
        return "", fmt.Errorf("failed to get management token, status code: %d", resp.StatusCode)
    }

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        log.Printf("Error decoding response: %v", err)
        return "", err
    }

    token, ok := result["access_token"].(string)
    if !ok {
        log.Printf("No access token in response: %v", result)
        return "", fmt.Errorf("no access token in response")
    }

    log.Printf("Successfully obtained management token")
    return token, nil
}
