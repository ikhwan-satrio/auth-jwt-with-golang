# GEMINI.md

## Project Overview

**auth-golang** is a Go-based authentication service providing user registration and login functionality. It follows a clean architecture with separated concerns for handlers, services, and configuration.

- **Primary Technologies:**
    - **Language:** Go 1.26.1
    - **Router:** [chi v5](https://github.com/go-chi/chi)
    - **ORM:** [GORM v1.31.1](https://gorm.io) (utilizing a generic `gorm.G` interface)
    - **Database:** SQLite (support for PostgreSQL and MySQL via configuration)
    - **Security:** JWT (`golang-jwt/jwt/v5`) and password hashing with `bcrypt`.
    - **Development:** [Air](https://github.com/cosmtrek/air) for live reloading.

---

## Project Structure

- `cmd/main.go`: Entry point for the application. Initializes config, database, and routes.
- `app/config/`:
    - `config.go`: Configuration management for Database and HTTP settings.
- `app/db/`:
    - `models.go`: GORM models.
    - `services.go`: Database connection and migration logic.
- `app/handler/`:
    - `auth_handler.go`: HTTP handlers for registration and login.
- `app/middlewares/`:
    - `auth_middleware.go`: JWT authentication middleware.
- `app/services/`:
    - `tokens/`: JWT token lifecycle management.

---

## Building and Running

### Configuration

The application can be configured via environment variables:
- `DB_DRIVER`: `sqlite`, `postgres`, or `mysql` (default: `sqlite`).
- `DB_DSN`: Database connection string (default: `./database.sqlite`).
- `HTTP_PORT`: Port to listen on (default: `3000`).

### Development Mode (with Air)

To start the server with live reloading:

```bash
air
```

*Note: `.air.toml` has been updated to point to `cmd/main.go`.*

### Manual Build

```bash
go build -o auth-golang cmd/main.go
./auth-golang
```

---

## API Endpoints

- `GET /`: Health check.
- `POST /auth/register`: Register a new user.
- `POST /auth/login`: Login and receive a JWT token.
- `GET /users`: (Protected) Fetch user details (requires Bearer token).

---

## Development Conventions

- **Handlers:** All HTTP logic belongs in `app/handler`.
- **Middlewares:** Global and route-specific middlewares are in `app/middlewares`.
- **Config:** Use `app/config` to manage environment-based settings.
- **Database:** Use the generic `gorm.G` interface for database operations.
