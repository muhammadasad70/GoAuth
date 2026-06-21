# GoAuth

GoAuth is a production-style Authentication and Authorization Service built to provide secure user management for modern web and mobile applications.

The project serves as a practical learning platform for backend engineering using Go, PostgreSQL, Redis, JWT authentication, clean architecture, security best practices, and Docker.

---

## Problem Statement

Almost every application requires authentication and authorization.

Many beginner projects implement only basic login and signup functionality without considering:

* Session Management
* Refresh Tokens
* Role-Based Access Control
* Email Verification
* Password Recovery
* Rate Limiting
* Security Best Practices

GoAuth aims to solve this problem by providing:

* Secure Authentication
* JWT Authorization
* Refresh Token Management
* Session Management
* Email Verification
* Password Recovery
* Role-Based Access Control
* API Security

---

## Technology Stack

### Backend

* Go
* Gin Framework
* JWT Authentication
* bcrypt Password Hashing
* Middleware
* REST APIs

### Database

* PostgreSQL

### Cache & Session Management

* Redis

### DevOps

* Git
* GitHub
* Docker
* Docker Compose

### Email Services

* SMTP
* Mailtrap

---

## Core Features

### Authentication

* User Registration
* User Login
* User Logout
* JWT Authentication
* Protected Routes

### Authorization

* Role-Based Access Control (RBAC)
* Admin Routes
* User Permissions

### Session Management

* Refresh Tokens
* Redis Sessions
* Logout From Single Device
* Logout From All Devices

### Account Security

* Password Hashing
* Email Verification
* Forgot Password
* Reset Password
* Rate Limiting

### Backend Engineering

* Clean Architecture
* Modular Monolith Design
* Middleware
* Repository Pattern
* Service Layer

---

## Project Documentation

Detailed documentation can be found inside the `docs` folder:

* project-overview.md
* architecture.md
* roadmap.md
* database-schema.md
* api-endpoints.md
* authentication-flow.md
* security-notes.md
* development-journal.md

---

## Current Status

Project is currently in the planning and architecture phase.

Next milestone:

* PostgreSQL Setup
* Environment Configuration
* Database Connection
* Users Table Migration

---

## Purpose

This project combines two goals:

1. Learning production-level authentication and authorization systems.
2. Developing professional backend engineering skills using Go.

The goal is to build a reusable authentication service while gaining practical experience in security, APIs, databases, middleware, session management, and scalable backend architecture.
