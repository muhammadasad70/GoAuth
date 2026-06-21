# Authentication Flow

## Registration

User

↓

Register API

↓

Validate Request

↓

Hash Password

↓

Store User

↓

Generate Verification Token

↓

Send Verification Email

---

## Login

User

↓

Login API

↓

Verify Email

↓

Compare Password

↓

Generate Access Token

↓

Generate Refresh Token

↓

Store Session

↓

Return Tokens

---

## Protected Route

Request

↓

JWT Middleware

↓

Validate Token

↓

Extract User ID

↓

Allow Access

---

## Refresh Token

Request

↓

Validate Refresh Token

↓

Check Redis Session

↓

Generate New Access Token

↓

Return Access Token

---

## Password Reset

User

↓

Forgot Password

↓

Generate Reset Token

↓

Send Email

↓

Reset Password

↓

Update Password Hash

↓

Success
