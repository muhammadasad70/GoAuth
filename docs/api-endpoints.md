# API Endpoints

## Authentication

### Register

POST /api/v1/auth/register

### Login

POST /api/v1/auth/login

### Logout

POST /api/v1/auth/logout

### Refresh Token

POST /api/v1/auth/refresh

---

## User

### Get Profile

GET /api/v1/users/me

### Update Profile

PUT /api/v1/users/me

### Change Password

PUT /api/v1/users/change-password

---

## Email Verification

### Send Verification

POST /api/v1/auth/send-verification

### Verify Email

GET /api/v1/auth/verify-email

---

## Password Reset

### Forgot Password

POST /api/v1/auth/forgot-password

### Reset Password

POST /api/v1/auth/reset-password

---

## Admin

### Get Users

GET /api/v1/admin/users

### Block User

PUT /api/v1/admin/users/:id/block

### Unblock User

PUT /api/v1/admin/users/:id/unblock
