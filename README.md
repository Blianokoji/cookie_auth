# 🔐 Cookie-based authN in Go

> Tiny demo that sets, reads, and clears a single HttpOnly cookie named `auth` to simulate login.

---

**Highlights**

- 🟦 Go 1.x
- 🍪 HttpOnly cookie
- 🎓 Minimal, educational

## What this is

A compact, educational cookie-auth simulation in Go. The handlers establish a very simple flow: set a cookie on login, read it on profile, and clear it on logout.

## Endpoints

| Method | Path                   | Purpose                                       |
|:------:|------------------------|-----------------------------------------------|
| GET    | `/ping`                | Health check — returns `Pong`                 |
| GET    | `/login?user={name}`   | Sets `auth` cookie to the provided username   |
| GET    | `/profile`             | Reads `auth` cookie; 401 if missing           |
| GET    | `/logout`              | Clears the cookie (`MaxAge=-1`)               |

## Quick run

```powershell
go run main.go
```

Then open:

- http://localhost:8080/ping
- http://localhost:8080/login?user=blessen
- http://localhost:8080/profile
- http://localhost:8080/logout

## Notes

- No database or persistence — this is a demo.
- Cookie is HttpOnly. `Secure=false` for local testing; set to `true` when serving over TLS.

## What I implemented

- Handlers in `internals/handlers`: `Ping`, `Login`, `Profile`, `Logout`.
- Basic router in `main.go` using `http.HandleFunc` on port `:8080`.
