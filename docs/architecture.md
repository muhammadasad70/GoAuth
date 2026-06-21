# Architecture

## Architecture Type

GoAuth follows a **Modular Monolithic Architecture** with **Clean Architecture Principles**.

The application is deployed as a single service, but the codebase is organized into independent modules to improve maintainability, scalability, and readability.

Unlike a traditional monolith where all logic is mixed together, GoAuth separates responsibilities into dedicated layers and modules.

---

# High-Level Architecture

```text
Client
   │
   ▼
Routes
   │
   ▼
Middleware
   │
   ▼
Handler
   │
   ▼
Service
   │
   ▼
Repository
   │
   ▼
Database / Redis
```

---

# Request Flow

### Step 1: Client Request

A client sends an HTTP request.

Example:

```http
POST /api/v1/auth/login
```

---

### Step 2: Routes

The route determines which handler should process the request.

Example:

```go
router.POST("/auth/login", authHandler.Login)
```

---

### Step 3: Middleware

Middleware performs cross-cutting concerns before the request reaches the handler.

Examples:

* JWT Validation
* Role Verification
* Rate Limiting
* Request Logging

---

### Step 4: Handler Layer

The handler receives the request and performs:

* Request parsing
* Request validation
* Calling service methods
* Sending HTTP responses

The handler should not contain business logic.

Example:

```text
LoginHandler
RegisterHandler
RefreshTokenHandler
```

---

### Step 5: Service Layer

The service layer contains business logic.

Responsibilities:

* User registration
* Login verification
* Password hashing
* Token generation
* Email verification logic
* Session management

Example:

```text
AuthService
UserService
SessionService
EmailService
```

The service layer is considered the core of the application.

---

### Step 6: Repository Layer

Repositories communicate with data sources.

Responsibilities:

* Database queries
* Data retrieval
* Data persistence
* Redis operations

Example:

```text
UserRepository
SessionRepository
```

The repository layer should contain no business logic.

---

### Step 7: Database Layer

The database layer stores persistent application data.

---

## PostgreSQL

Used for:

* Users
* Email Verification Tokens
* Password Reset Tokens
* User Roles
* Account Status

### Tables

```text
users
email_verifications
password_resets
```

---

## Redis

Used for:

* Refresh Token Sessions
* Session Revocation
* Rate Limiting
* Temporary Authentication Data

Redis provides very fast access for authentication-related operations.

---

# Project Modules

## Auth Module

Responsibilities:

* Register User
* Login User
* Logout User
* Refresh Access Token

Files:

```text
internal/auth/
```

---

## User Module

Responsibilities:

* User Profile
* Update Profile
* Change Password

Files:

```text
internal/user/
```

---

## Session Module

Responsibilities:

* Session Creation
* Session Revocation
* Refresh Token Management

Files:

```text
internal/session/
```

---

## Email Module

Responsibilities:

* Verification Emails
* Password Reset Emails

Files:

```text
internal/email/
```

---

## Token Module

Responsibilities:

* JWT Generation
* JWT Validation
* Refresh Token Generation

Files:

```text
internal/token/
```

---

## Middleware Module

Responsibilities:

* Authentication Middleware
* Authorization Middleware
* Rate Limiting Middleware

Files:

```text
internal/middleware/
```

---

# Folder Structure

```text
go-auth-service/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── auth/
│   ├── user/
│   ├── session/
│   ├── email/
│   ├── middleware/
│   ├── token/
│   ├── config/
│   ├── database/
│   ├── routes/
│   └── utils/
│
├── migrations/
├── docs/
├── tests/
│
├── Dockerfile
├── docker-compose.yml
├── .env.example
└── README.md
```

---

# Why Modular Monolith?

This architecture was chosen because:

* Easier to understand for a junior backend developer
* Faster development compared to microservices
* Easier debugging
* Simpler deployment
* Suitable for small and medium-sized systems
* Can later be split into microservices if needed

---

# Benefits

* Clear separation of concerns
* Easy maintenance
* Better testability
* Scalable code organization
* Industry-standard backend structure
* Portfolio-ready project architecture

---

# Future Evolution

As the system grows, individual modules can be extracted into separate microservices.

Potential future services:

* Authentication Service
* Notification Service
* Audit Service
* User Service

For Version 1, a Modular Monolith provides the best balance between simplicity and production-level architecture.
