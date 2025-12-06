# Go Architecture Template

A clean, reusable Go project template following modern architectural practices.
This structure provides clear separation of concerns, scalability, and maintainability for real-world backend services.

---

## Project Structure

```sql
go-architecture-template/
│
├── cmd/
│   └── app/
│       └── main.go          # Application entrypoint
│
├── internal/
│   ├── api/                 # HTTP handlers and response utilities
│   │   ├── handler.go
│   │   └── response/
│   │       └── response.go
│   │
│   ├── domain/              # Pure domain models (business entities)
│   │   └── example.go
│   │
│   ├── service/             # Business logic / use cases
│   │   └── example_service.go
│   │
│   ├── storage/             # Persistence layer (DB, memory, etc.)
│   │   └── postgres/
│   │       └── example_repo.go
│   │
│   ├── config/              # Application configuration
│   │   └── config.go
│   │
│   └── util/                # Utilities (logger, helpers, etc.)
│       └── logger.go
│
├── pkg/                     # Optional public packages (exportable)
│
├── go.mod
└── README.md
```
---

## Goals of This Template

- Provide a clean starting point for scalable Go services
- Encourage layered architecture and dependency boundaries
- Keep business logic independent from infrastructure
- Avoid framework lock-in
- Keep the template minimal and easy to extend

---

## How to Use This Template

### 1. Clone this template

```bash
git clone https://github.com//go-architecture-template
```

### 2. Copy it for a new project

```bash
rsync -a go-architecture-template/  your-new-service
```

### 3. Replace the module name

```bash
go mod edit -module=your-new-service
go mod tidy
```

---

### 4. Start implementing your domain, service, and storage logic

- Add real domain models under internal/domain
- Add business logic under internal/service
- Add database implementations in internal/storage
- Register your routes in handler.go

---

## Architecture Principles Used

- Separation of concerns
- Dependency inversion
- Clear layering (API → Service → Storage)
- No circular dependencies
- Domain-driven mindset (but lightweight)
- Testability and isolation

---

## License

MIT License — feel free to use and adapt.
