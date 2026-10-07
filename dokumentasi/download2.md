# Package yang wajib di download

go get github.com/redis/go-redis/v9

# Framework HTTP (Fiber v3)

go get github.com/gofiber/fiber/v3

# Database Driver

go get github.com/go-sql-driver/mysql

# Autentikasi & Keamanan

go get github.com/golang-jwt/jwt/v5
go get github.com/alexedwards/argon2id
go get golang.org/x/crypto

# Validasi Data & Utility

go get github.com/go-playground/validator/v10
go get github.com/google/uuid

# Logging & Environment Config

go get github.com/sirupsen/logrus
go get github.com/joho/godotenv

# Dependency Injection

go get github.com/google/wire

# Email

go get gopkg.in/gomail.v2

# Testing

go get github.com/stretchr/testify

# Bersihkan & sinkronkan go.mod

go mod tidy

# Download semuanya sekaligus

go get github.com/gofiber/fiber/v3
github.com/go-sql-driver/mysql
github.com/golang-jwt/jwt/v5
github.com/alexedwards/argon2id
golang.org/x/crypto
github.com/go-playground/validator/v10
github.com/google/uuid
github.com/sirupsen/logrus
github.com/joho/godotenv
github.com/google/wire
gopkg.in/gomail.v2
github.com/stretchr/testify

go get github.com/redis/go-redis/v9&& go mod tidy
