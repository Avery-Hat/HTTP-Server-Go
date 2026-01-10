# HTTP-Server-Go

Boot.dev — Learn HTTP Servers in Go (up through Chapter 7)

A simple REST API written in Go for a Twitter-like application called **Chirpy**.  
Built as part of the Boot.dev curriculum, this project focuses on:

- HTTP servers in Go
- Authentication & authorization
- JWT access tokens and refresh tokens
- PostgreSQL with SQLC
- Secure password handling
- Webhooks & API key verification

---

## Features

- User registration and login
- JWT-based authentication
- Refresh token rotation & revocation
- Create, read, filter, sort, and delete chirps
- Authorization: users can only modify their own data
- Chirpy Red membership via secure webhooks
- Admin metrics & reset endpoints (dev-only)

---

## API Endpoints

### Auth & Users
- `POST /api/users` — create a user
- `PUT /api/users` — update authenticated user's email/password
- `POST /api/login` — login and receive access + refresh tokens
- `POST /api/refresh` — exchange refresh token for a new access token
- `POST /api/revoke` — revoke a refresh token

### Chirps
- `GET /api/chirps`
  - Optional query params:
    - `author_id=<uuid>`
    - `sort=asc|desc`
- `GET /api/chirps/{chirpID}`
- `POST /api/chirps` *(authenticated)*
- `DELETE /api/chirps/{chirpID}` *(authenticated, owner only)*

### Webhooks
- `POST /api/polka/webhooks` — Polka webhook to upgrade users to Chirpy Red (API key protected)

### Health & Admin
- `GET /api/healthz`
- `GET /admin/metrics`
- `POST /admin/reset` *(dev only)*

---

## Environment Variables

Create a `.env` file (do **not** commit it):

```env
DB_URL=postgres://user:password@localhost:5432/chirpy?sslmode=disable
JWT_SECRET=your-secret-key
POLKA_KEY=your-polka-api-key
PLATFORM=dev
```
---

## Running Locally
```md
set -a
source .env
set +a

goose postgres "$DB_URL" up
go run .
```