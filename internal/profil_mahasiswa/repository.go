package profilmahasiswa

import (
	"context"
)

// Interface untuk postgreSQL

type MhsRepo interface {
	TambahMhs(ctx context.Context, domain Profil) (Profil, error)
	GetById(ctx context.Context, id int) (Profil, error)
	GetByEmailUsername(ctx context.Context, email, username string) (Profil, error)
}

// Interface untuk Redis
type RedisUser interface {
	CreateUser(ctx context.Context, do Profil, key string) string
	GetUserEmail(ctx context.Context, key string) string

	SimpanOTP(ctx context.Context, otp string, keyUser string) string // menambahkan otp sekaligus mengembalikan otp nya, redis string
	GetUser(ctx context.Context, key string) Profil
	GetOTP(ctx context.Context, key string) DataOTP
}
