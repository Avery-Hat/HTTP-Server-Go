# HTTP-Server-Go
Boot.Dev lesson part 18: Learn HTTP Servers in Go. This is up to the end of Chapter 7.

A simple REST API written in Go for a Twitter-like application called **Chirpy**.  
This project is built as part of the Boot.dev curriculum and focuses on:

- HTTP servers in Go
- Authentication & authorization
- JWT access tokens and refresh tokens
- PostgreSQL with SQLC
- Secure password handling

---

## Features

- User registration and login
- JWT-based authentication
- Refresh token rotation & revocation
- Create, read, and delete chirps
- Authorization: users can only modify their own data
- Admin metrics & reset endpoints (dev-only)

---

## API Endpoints

### Auth
- `POST /api/users` — create a user
- `POST /api/login` — login and receive access + refresh tokens
- `POST /api/refresh` — exchange refresh token for a new access token
- `POST /api/revoke` — revoke a refresh token
- `PUT /api/users` — update authenticated user's email/password

### Chirps
- `GET /api/chirps`
- `GET /api/chirps/{chirpID}`
- `POST /api/chirps` *(authenticated)*
- `DELETE /api/chirps/{chirpID}` *(authenticated, owner only)*

### Health & Admin
- `GET /api/healthz`
- `GET /admin/metrics`
- `POST /admin/reset` *(dev only)*

---

## Environment Variables

Create a `.env` file:

```env
DB_URL=postgres://user:password@localhost:5432/chirpy?sslmode=disable
JWT_SECRET=your-secret-key
PLATFORM=dev
