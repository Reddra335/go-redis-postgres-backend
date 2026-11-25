package profilmahasiswa

import (
	"fmt"
	"newProjek/internal/shared/logger"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/sirupsen/logrus"
)

type ProfilHandler interface {
	Register(req fiber.Ctx) error
	VerifyOTP(req fiber.Ctx) error
	ClearCookie(c fiber.Ctx) error
}

type profilHandlerImpln struct {
	service ProfileService
}

func NewProfilHandler(service ProfileService) ProfilHandler {

	return &profilHandlerImpln{

		service: service,
	}

}

type Resp struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Prodi    string `json:"prodi"`
}

func (h *profilHandlerImpln) Register(c fiber.Ctx) error {
	log := logger.New()
	var req ReqMahasiswa
	err := c.Bind().Body(&req)
	if err != nil {
		log.WithError(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  "BAD REQUEST",
			"message": "REQUEST tidak sesuai",
		})
	}
	redisotp := c.Cookies("redisKey")
	ck := fmt.Sprintf("cookie masih ada: %s", redisotp)
	log.Info(ck)
	respone := h.service.Register(c.Context(), req, redisotp)
	if respone.Email != "" {
		return c.Status(302).JSON(fiber.Map{
			"code":    302,
			"status":  "Move To Verify",
			"message": "Lanjutkan Ke Verify",
		})
	}
	ambil := &Resp{
		Name:     respone.Name,
		Username: respone.Username,
		Email:    respone.Email,
		Prodi:    respone.Prodi,
	}
	if respone.KeyUser != "" {
		dm := fmt.Sprintf("key dari key response anda adalah %s", respone.KeyUser)
		log.Info(dm)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "redisKey",
		Value:    respone.KeyUser,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
		MaxAge:   int(30 * time.Minute.Seconds()),
	})
	redisotptest := c.Cookies("redisKey")
	d := fmt.Sprintf("key dari cookie anda adalah %s", redisotptest)
	log.Info(d)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    201,
		"status":  "Created Sementara",
		"message": ambil,
	})
}
func (h *profilHandlerImpln) VerifyOTP(c fiber.Ctx) error {
	rediskey := c.Cookies("redisKey")
	if rediskey == "" {
		log.Error("redis key dari cookie kosong")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  "BAD REQUEST",
			"message": "Cookie kosong",
		})

	}
	type ReqOtp struct {
		OTP string `json:"otp"`
	}
	var req ReqOtp
	log := logger.New()
	err := c.Bind().Body(&req)

	if err != nil {

		log.WithFields(logrus.Fields{

			"operation": "verify",
			"layer":     "handler",
		}).WithError(err).Error("ada yang salah dengan req")

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  "BAD REQUEST",
			"message": "Request salah",
		})
	}

	response, err := h.service.VerifyOTP(c.Context(), rediskey, req.OTP)

	if err != nil {

		log.WithFields(logrus.Fields{

			"operation": "verify",
			"layer":     "handler",
		}).WithError(err).Error("ada yang salah dengan verify")

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"status":  "BAD REQUEST",
			"message": err,
		})
	}

	ambil := &Resp{
		Name:     response.Name,
		Username: response.Username,
		Email:    response.Email,
		Prodi:    response.Prodi,
	}

	c.ClearCookie("redisKey")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"status":  "Created",
		"message": ambil,
	})

}

func (h *profilHandlerImpln) ClearCookie(c fiber.Ctx) error {
	a := c.Cookies("redisKey")

	log.Info(fmt.Sprintf("isi dari Key anda: %s", a))

	return c.JSON(fiber.Map{
		"message": "ambil cookie",
	})
}
