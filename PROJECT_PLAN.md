# Titan - Go SaaS Backend Template

## Vision
A reusable, production-ready Go backend that can be copy-pasted to any client project. Ship faster than competitors.

---

## Tech Stack
- **Language:** Go 1.22+
- **Router:** Standard library (`net/http`)
- **Database:** PostgreSQL (via `pgx`)
- **Auth:** JWT (access + refresh tokens)
- **Payments:** Stripe (later)
- **Deployment:** Docker + Cloud Run

---

## Phases

### Phase 1: Foundation ✅
- [x] Project structure (`cmd/`, `internal/`)
- [x] Configuration (CLI flags)
- [x] Structured logging (`slog`)
- [x] Database connection pool
- [x] JSON helpers (`readJSON`, `writeJSON`)
- [x] Centralized error responses
- [x] Users CRUD with repository pattern
- [x] Optimistic locking

### Phase 2: Production Hardening
- [ ] Migrations (`golang-migrate`)
- [ ] Middleware (logging, panic recovery, rate limiting)
- [ ] Input validation
- [ ] Graceful shutdown
- [ ] CORS handling

### Phase 3: Authentication
- [ ] Password hashing (bcrypt)
- [ ] User registration & login
- [ ] JWT access tokens
- [ ] Refresh token rotation
- [ ] Protected routes middleware

### Phase 4: Authorization
- [ ] Role-based access control (RBAC)
- [ ] Permission system
- [ ] Admin vs User routes

### Phase 5: Production Features
- [ ] Email sending (welcome, password reset)
- [ ] Background jobs
- [ ] API versioning
- [ ] Swagger/OpenAPI docs
- [ ] Health checks (deep)

### Phase 6: Payments (Optional)
- [ ] Stripe customer creation
- [ ] Subscription management
- [ ] Webhook handling
- [ ] Billing portal

---

## Current Status
**Phase 1 Complete** - Ready for Phase 2 (Migrations & Middleware)

---

## File Structure (Target)