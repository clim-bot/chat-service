# Chat Service

A simple chat application built with Golang, Gin, WebSockets, Gorm, PostgreSQL, and Auth0 for authentication.

## Features

- User authentication using Auth0
- Real-time messaging using WebSockets
- Support for private and group chats

## Requirements

- Go 1.20 or higher
- PostgreSQL
- Auth0 account
- Docker
- Docker Compose
- Make

## Folder Structure
```go
chat-service/
├── controllers/
│   ├── auth.go
│   ├── chat.go
├── models/
│   ├── message.go
│   ├── user.go
├── routes/
│   ├── auth.go
│   ├── chat.go
├── utils/
│   ├── database.go
│   ├── websocket.go
├── main.go
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
├── go.sum
├── .env
```

## Setup

### Auth0 Configuration

1. **Create an Auth0 Account**

   Sign up at [Auth0](https://auth0.com/).

2. **Create a New Application**

   - Log in to your Auth0 dashboard.
   - Navigate to the "Applications" section from the left sidebar.
   - Click on the "Create Application" button.
   - Choose a name for your application.
   - Select the type of application: "Regular Web Applications".
   - Click "Create".

3. **Configure the Application**

   - In the settings page, copy the `Client ID` and `Client Secret`.
   - Set the "Allowed Callback URLs" to `http://localhost:8080/auth/callback`.
   - Set the "Allowed Logout URLs" to `http://localhost:8080`.
   - Set the "Allowed Web Origins" to `http://localhost:8080`.

4. **Create a Machine-to-Machine Application**

   - Navigate to the "Applications" section.
   - Click on the "Create Application" button.
   - Choose a name for your application.
   - Select "Machine to Machine Applications".
   - Click "Create".

5. **Set Up API Permissions**

   - In the "APIs" section, select the Auth0 Management API.
   - Grant the application the necessary permissions (e.g., `create:users`).

6. **Get the Management API Token**

   - Go to the "API Explorer" tab in the Auth0 Management API section.
   - Click on "Copy Token" to get a temporary token for testing.

### Project Setup

1. **Clone the repository**

    ```bash
    git clone https://github.com/clim-bot/chat-service.git
    cd chat-service
    ```

2. **Create a `.env` file**

    ```env
    AUTH0_DOMAIN=your-tenant.auth0.com
    AUTH0_CLIENT_ID=your-client-id
    AUTH0_CLIENT_SECRET=your-client-secret
    DATABASE_URL=postgres://username:password@localhost:5432/chat_db?sslmode=disable
    SESSION_SECRET=your-session-secret
    ```

3. **Install dependencies**

    ```bash
    go mod tidy
    ```

4. **Run the application**

    ```bash
    go run main.go
    ```

## API Endpoints

### Authentication

##### Register

```http
POST /auth/register
```
Registers a new user using Auth0.

#### Login
```http
GET /auth/login
```
Redirects to Auth0 for user authentication.

### Callback
```http
GET /auth/callback
```
Callback endpoint for Auth0 authentication.

### Logout
```http
GET /auth/logout
```
Logs the user out and clears the session.

### Chat
WebSocket Connection
```http
GET /chat/ws
```
Establishes a WebSocket connection for real-time messaging.

1. Login

Open a browser and navigate to `http://localhost:8080/auth/login` to log in using Auth0.

2. Connect to WebSocket

You can use a tool like websocat to test the WebSocket connection:
```bash
websocat ws://localhost:8080/chat/ws
```

Or use curl to test other HTTP endpoints:
```bash
curl -i http://localhost:8080/auth/login
curl -i http://localhost:8080/auth/logout
```

## License
This project is licensed under the MIT License.