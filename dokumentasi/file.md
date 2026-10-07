tamplate_golang_docker_redist/
├── cmd/
│ ├── main.go
│ ├── providers.go
│ ├── router.go
│ ├── wire.go
│ └── wire_gen.go
├── internal/
│ ├── helperdb/
│ │ └── executor.go
│ ├── shared/
│ │ ├── auth/ # JWT Token Generator & Validator
│ │ ├── config/ # Inisialisasi DB & Redis
│ │ │ ├── database.go
│ │ │ └── redis.go # [BARU] Client Redis
│ │ ├── logger/ # Logrus Config
│ │ ├── middleware/ # Fiber v3 JWT Middleware & Rate Limiter
│ │ ├── pass/ # Argon2id Hashing
│ │ └── send_email/ # Gomail & OTP Generator
│ └── user/ # Domain User (Dipersimpel)
│ ├── entity.go # Struct Database Domain
│ ├── dto.go # Struct Request & Response API
│ ├── repository.go # Interface & Impl MySQL + Redis
│ ├── service.go # Interface & Impl Business Logic
│ ├── handler.go # Fiber v3 HTTP Handlers
│ └── router.go # Route endpoints
├── .env
├── .env.example
├── Dockerfile # [BARU] Multi-stage build Go
├── docker-compose.yml # [BARU] Orchestration Go + MySQL + Redis
├── go.mod
└── go.sum
