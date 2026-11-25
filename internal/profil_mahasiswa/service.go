package profilmahasiswa

import (
	"context"
	"errors"
	"fmt"
	generatecode "newProjek/internal/shared/helper/generate_code"
	generateenkripsi "newProjek/internal/shared/helper/generate_enkripsi"
	"newProjek/internal/shared/logger"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

// Interface
type ProfileService interface {
	Register(ctx context.Context, req ReqMahasiswa, keyvalid string) ResponseMahasiswa
	VerifyOTP(ctx context.Context, key string, reqOtp string) (ResponseMahasiswa, error)
	LoginService(ctx context.Context, username, password, email string)
}

// send email
type EmailSender interface {
	SendOTP(emailUser string, otp string) error
}

type profileServiceImpln struct {
	valid *validator.Validate
	repo  MhsRepo
	rds   RedisUser
	send  EmailSender
}

func NewProfileServiceImpln(repo MhsRepo, valid *validator.Validate, Rds RedisUser, send EmailSender) ProfileService {

	return &profileServiceImpln{
		valid: valid,
		repo:  repo,
		rds:   Rds,
		send:  send,
	}

}
func ToDomainProfil(req ReqMahasiswa) Profil {
	da := Profil{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Prodi:    req.Prodi,
		Email:    req.Email,
	}
	return da
}

func (s *profileServiceImpln) Register(ctx context.Context, req ReqMahasiswa, keyvalid string) ResponseMahasiswa {
	log := logger.New()

	// if keyvalid != "" {

	// 	log.WithFields(logrus.Fields{
	// 		"operation": "Register",
	// 		"layer":     "service",
	// 	}).Error("silahkan lanjut ke verify")
	// 	return ResponseMahasiswa{}

	// }
	err := s.valid.Struct(req)
	if err != nil {
		log.WithFields(logrus.Fields{

			"operation": "Register",
			"layer":     "Service Mahasiswa",
		}).WithError(err).Error("ada yang salah dengan validate")

		return ResponseMahasiswa{}
	}
	profil := ToDomainProfil(req)

	idRedis := s.rds.CreateUser(ctx, profil, generateenkripsi.IdRds())
	if idRedis == "" {

		log.WithFields(logrus.Fields{

			"operation": "Register",
			"layer":     "Service Mahasiswa",
		}).Error("id redis kosong:", idRedis)

		return ResponseMahasiswa{}

	}

	email := s.rds.GetUserEmail(ctx, idRedis)

	if email == "" {

		log.WithFields(logrus.Fields{

			"operation": "Register",
			"layer":     "Service Mahasiswa",
		}).WithError(err).Error("email anda kosong")

		return ResponseMahasiswa{}

	}

	keyUser := "otp:" + idRedis
	msg := fmt.Sprintf("key otp anda %s", keyUser)
	log.Info(msg)

	otp, err := generatecode.GenerateOTP()
	if err != nil {

		log.WithFields(logrus.Fields{

			"operation": "Register",
			"layer":     "Service Mahasiswa",
		}).WithError(err).Error("ada yang salah dengan generate otp")
		return ResponseMahasiswa{}

	}

	err = s.send.SendOTP(email, otp)
	if err != nil {
		log.WithFields(logrus.Fields{

			"operation": "Register",
			"layer":     "Service Mahasiswa",
		}).WithError(err).Error("gomail bermasalah")

	}

	key := s.rds.SimpanOTP(ctx, otp, idRedis)

	return ResponseMahasiswa{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Prodi:    req.Prodi,
		KeyUser:  key,
	}

}

func (s *profileServiceImpln) VerifyOTP(ctx context.Context, key string, reqOtp string) (ResponseMahasiswa, error) {

	log := logger.New()
	errReq := s.valid.Var(reqOtp, "required,min=6,max=6")
	if errReq != nil {
		log.WithError(errReq).Error("otp tidak boleh kosong")
		return ResponseMahasiswa{}, errors.New("OTP anda salah atau kosong")

	}

	if key == "" {
		log.WithFields(logrus.Fields{
			"operation": "VerifyOTP",
			"layer":     "service",
		}).Error("key kosong\n")
		return ResponseMahasiswa{}, errors.New("key kosong")
	}
	errKey := s.valid.Var(key, "required")
	if errKey != nil {
		log.WithError(errReq).Error("Key tidak boleh kosong")
		return ResponseMahasiswa{}, errors.New("Key anda salah atau kosong")

	}
	data := s.rds.GetOTP(ctx, key)
	if data.KeyUser == "" {

		return ResponseMahasiswa{}, errors.New("GetOTP redis kosong")
	}

	if data.OTP != reqOtp {
		log.WithFields(logrus.Fields{
			"operation": "comper",
			"layer":     "service",
		}).Error("OTP Salah")
		return ResponseMahasiswa{}, errors.New("otp salah")
	}
	log.Info(fmt.Sprintf("isi func  GETotp response key %s", data.KeyUser))

	dataMhs := s.rds.GetUser(ctx, data.KeyUser)

	if dataMhs.Username == "" {
		log.WithFields(logrus.Fields{

			"operation": "verify otp",
			"layer":     "service",
		}).Error("data kosong")

		return ResponseMahasiswa{}, errors.New("verify gagal")
	}

	mhs, err := s.repo.TambahMhs(ctx, dataMhs)
	if err != nil {

		log.WithFields(logrus.Fields{
			"operation": "create",
			"layer":     "service",
		}).WithError(err).Error("data gagal di tambahkan")

		return ResponseMahasiswa{}, errors.New("service gagal")
	}

	log.Info("Data berhasil di pindahkan")
	return ResponseMahasiswa{
		Name:     mhs.Name,
		Username: mhs.Username,
		Email:    mhs.Email,
		Prodi:    mhs.Prodi,
	}, nil
}

func (s *profileServiceImpln) LoginService(ctx context.Context, username, password, email string) {

s.repo.GetById(ctx,)








}
