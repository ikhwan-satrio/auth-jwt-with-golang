# Auth Golang Example (JWT)

A robust authentication service implemented in Go, featuring user registration, login, and protected routes using JWT (JSON Web Tokens). This project follows a clean architecture pattern and uses GORM for database operations.

## 🚀 Features

- **User Registration**: Secure password hashing with `bcrypt`.
- **User Login**: Identity verification and JWT issuance.
- **JWT Authentication**: Middleware-protected routes.
- **Clean Architecture**: Separation of concerns between handlers, services, and models.
- **Multi-Database Support**: Support for SQLite, PostgreSQL, and MySQL.
- **Live Reload**: Integrated with [Air](https://github.com/cosmtrek/air) for an efficient development workflow.

## 🛠️ Tech Stack

- **Language**: [Go](https://go.dev/) (v1.26.1)
- **HTTP Router**: [Chi v5](https://github.com/go-chi/chi)
- **ORM**: [GORM](https://gorm.io/)
- **Security**: [golang-jwt/jwt](https://github.com/golang-jwt/jwt) & `bcrypt`
- **Database**: SQLite (default)

## 📁 Project Structure

```text
├── app/
│   ├── config/       # Configuration management
│   ├── db/           # Database models and connection logic
│   ├── handlers/     # HTTP request handlers
│   ├── middlewares/  # JWT & other middlewares
│   ├── routes/       # Route definitions
│   └── services/     # Business logic (Auth, Tokens)
├── cmd/
│   └── main.go       # Application entry point
├── .air.toml         # Air configuration for live reload
└── go.mod            # Go modules and dependencies
```

## ⚙️ Configuration

The application can be configured using environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_DRIVER` | Database engine (`sqlite`, `postgres`, `mysql`) | `sqlite` |
| `DB_DSN` | Database connection string | `./database.sqlite` |
| `HTTP_PORT` | Port for the HTTP server | `3000` |

## 🏃 Getting Started

### Prerequisites

- Go 1.26.1 or later
- (Optional) [Air](https://github.com/cosmtrek/air) for live reloading

### Installation

1. Clone the repository:
   ```bash
   git clone <your-repo-url>
   cd auth-golang
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Application

**Development Mode (with Air):**
```bash
air
```

**Manual Build & Run:**
```bash
go build -o auth-golang cmd/main.go
./auth-golang
```

## 🛣️ API Endpoints

### Public Endpoints

#### 1. Health Check
- **URL**: `GET /`
- **Response**: `auth-golang service is running`

#### 2. Register
- **URL**: `POST /auth/register`
- **Body**:
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword123"
  }
  ```
- **Success Response**: `201 Created`
  ```json
  {
    "message": "register complete"
  }
  ```

#### 3. Login
- **URL**: `POST /auth/login`
- **Body**:
  ```json
  {
    "email": "john@example.com",
    "password": "securepassword123"
  }
  ```
- **Success Response**: `200 OK`
  ```json
  {
    "token": "your.jwt.token.here"
  }
  ```

### Protected Endpoints
*Requires Header: `Authorization: Bearer <token>`*

#### 1. Get Users
- **URL**: `GET /users`
- **Response**: List of registered users.

## 🔒 Security Notes

- Passwords are never stored in plain text; they are hashed using `bcrypt`.
- JWT tokens should be stored securely on the client side (e.g., HttpOnly cookies).
- Ensure the `DB_DSN` and any secret keys are managed securely in production environments.

## 📄 License

[MIT](LICENSE) (or your preferred license)
