package middleware

import (
	"database/sql"
	profilmahasiswa "newProjek/internal/Profil_Mahasiswa"
	"newProjek/internal/shared/auth"

	"github.com/gofiber/fiber/v3"
)

func JWTMiddleware(userRepo repositoryUser.UsersRepo, db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		// 1. Ambil header Authorization
		authHeader := c.Cookies("")

		// 2. Cek apakah header ada dan berawalan "Bearer "
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    401,
				"status":  "UNAUTHORIZED",
				"message": "Token tidak ditemukan atau format salah",
			})
		}

		// 4. Validasi token (cek signature, algoritma, dan expired time)
		claims, err := auth.ValidateJwt(authHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    401,
				"status":  "UNAUTHORIZED",
				"message": "Token tidak valid atau sudah kedaluwarsa",
			})
		}

		// 5. Verifikasi pengirim token (cek apakah user masih ada di database)
		// Menggunakan c.UserContext() yang merupakan standar Go context.Context

	profilmahasiswa.MhsRepo.GetById()
		if err != nil || user.ID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    401,
				"status":  "UNAUTHORIZED",
				"message": "Akun user tidak ditemukan atau sudah tidak aktif",
			})
		}

		// 6. Simpan claims ke context Fiber untuk diakses di Handler
		c.Locals("user", claims)

		// 7. Lanjut ke handler
		return c.Next()
	}
}
