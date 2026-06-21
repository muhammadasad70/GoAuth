# Security Notes

## Password Security

Passwords are never stored directly.

bcrypt is used to generate password hashes.

---

## JWT Security

Access Tokens:

* Short Expiration
* Stateless Authentication

Refresh Tokens:

* Stored in Redis
* Revocable

---

## Email Verification

Users must verify email before accessing protected features.

---

## Password Reset

Reset links use one-time tokens.

Tokens expire automatically.

---

## Session Management

Refresh Tokens are stored as hashed values.

Sessions can be revoked.

---

## Rate Limiting

Prevent brute-force attacks.

Applied on:

* Login Endpoint
* Register Endpoint
* Password Reset Endpoint

---

## Validation

All user inputs must be validated.

Examples:

* Email Format
* Password Length
* Required Fields

---

## Future Security Improvements

* Google OAuth
* GitHub OAuth
* MFA
* Device Fingerprinting
* Audit Logs
