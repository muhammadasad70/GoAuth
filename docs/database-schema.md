# Database Schema

## users

| Column            | Type           |
| ----------------- | -------------- |
| id                | UUID           |
| name              | VARCHAR        |
| email             | VARCHAR UNIQUE |
| password_hash     | TEXT           |
| role              | VARCHAR        |
| is_email_verified | BOOLEAN        |
| is_active         | BOOLEAN        |
| created_at        | TIMESTAMP      |
| updated_at        | TIMESTAMP      |

## sessions

| Column             | Type      |
| ------------------ | --------- |
| id                 | UUID      |
| user_id            | UUID      |
| refresh_token_hash | TEXT      |
| device_info        | TEXT      |
| ip_address         | TEXT      |
| expires_at         | TIMESTAMP |
| revoked_at         | TIMESTAMP |
| created_at         | TIMESTAMP |

## email_verifications

| Column     | Type      |
| ---------- | --------- |
| id         | UUID      |
| user_id    | UUID      |
| token_hash | TEXT      |
| expires_at | TIMESTAMP |
| used_at    | TIMESTAMP |
| created_at | TIMESTAMP |

## password_resets

| Column     | Type      |
| ---------- | --------- |
| id         | UUID      |
| user_id    | UUID      |
| token_hash | TEXT      |
| expires_at | TIMESTAMP |
| used_at    | TIMESTAMP |
| created_at | TIMESTAMP |

## Relationships

users

↓

sessions

↓

email_verifications

↓

password_resets
