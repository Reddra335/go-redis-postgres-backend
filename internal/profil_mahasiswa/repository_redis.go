package profilmahasiswa

import (
	"context"
	"encoding/json/v2"
	"errors"
	"fmt"
	"newProjek/internal/shared/logger"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type implnRedisUser struct {
	Rdb *redis.Client
}

func NewImplnRedisUser(rdb *redis.Client) RedisUser {

	return &implnRedisUser{Rdb: rdb}
}

//SELECT id, name, username, email, password, prodi FROM mahasiswa.profil WHERE id=2;

// Create User
func (r *implnRedisUser) CreateUser(ctx context.Context, do Profil, key string) string {

	log := logger.New()
	if key == "" {

		log.WithFields(logrus.Fields{
			"operation": "Create User Redis",
			"layer":     "Redis",
		}).Error("key kosong")
		return ""
	}
	draft := UserDraft{

		ID:       do.ID,
		Name:     do.Name,
		Username: do.Username,
		Email:    do.Email,
		Password: do.Password,
		Prodi:    do.Prodi,
	}
	err := r.Rdb.HSet(ctx, key, map[string]any{
		"id":       draft.ID,
		"name":     draft.Name,
		"username": draft.Username,
		"email":    draft.Email,
		"password": draft.Password,
		"prodi":    draft.Prodi,
	}).Err()
	if err != nil {

		log.WithFields(logrus.Fields{
			"operation": "Create User Redis",
			"layer":     "Redis",
		}).WithError(err).Error(err)
		return ""
	}

	errX := r.Rdb.Expire(ctx, key, 30*time.Minute).Err()

	if errX != nil {
		log.WithFields(logrus.Fields{

			"operation": "Create User Redis",
			"layer":     "Redis",
		}).WithError(errX).Fatal("Terjadi kesalahan dalam expire hash")
		return ""
	}

	log.WithFields(logrus.Fields{
		"operation": "Create User Redis",
		"layer":     "Redis",
	}).Info("Data berhasil dibuat")
	return key

}

// Get User Email

func (r *implnRedisUser) GetUserEmail(ctx context.Context, key string) string {
	log := logger.New()
	if key == "" {

		err := errors.New("Key Anda bermasalah")
		log.WithFields(logrus.Fields{
			"operation": "Get Email",
			"layer":     "Redis",
		}).Error(err)
		return ""
	}
	email, err := r.Rdb.HGet(ctx, key, "email").Result()
	if err != nil {
		log.WithFields(logrus.Fields{
			"operation": "Get Email",
			"layer":     "Redis",
		}).WithError(err).Error("get email redis bermasalah")
		return ""
	}

	log.WithFields(logrus.Fields{
		"operation": "Get Email",
		"layer":     "Redis",
	}).Info("Data berhasil diambil email anda", email)
	return email
}

// Simpan OTP dan ambil
func (r *implnRedisUser) SimpanOTP(ctx context.Context, otp string, keyUser string) string {
	log := logger.New()

	if otp == "" || keyUser == "" {

		log.WithFields(logrus.Fields{
			"operation": "create otp",
			"layer":     "Redis",
		}).Error("ada yang kosong di simpan otp")
		return ""
	}

	data := map[string]any{

		"keyuser": keyUser,
		"otp":     otp,
	}

	dataJ, errJ := json.Marshal(data)

	if errJ != nil {
		log.WithFields(logrus.Fields{
			"operation": "json redis",
			"layer":     "redis",
		}).WithError(errJ).Error("gagal mengubah ke json")
		return ""

	}

	dataOTP := string(dataJ)
	key := "otp:" + keyUser
	//Simpan OTP
	err := r.Rdb.Set(ctx, key, dataOTP, 5*time.Minute).Err()
	if err != nil {
		log.WithFields(logrus.Fields{

			"operation": "create otp",
			"layer":     "Redis",
		}).WithError(err).Error("ada yang salah dengan create otp")
		return ""
	}

	log.Info("OTP berhasil di simpan")

	return key
}

// Ambil User
func (s *implnRedisUser) GetUser(ctx context.Context, key string) Profil {
	log := logger.New()
	if key == "" {

		errK := errors.New("key kosong")
		log.WithFields(logrus.Fields{
			"operation": "Ambil User",
			"key":       "Redis",
		}).WithError(errK).Error("key anda kosong")

		return Profil{}
	}

	userRes, err := s.Rdb.HGetAll(ctx, key).Result()

	if err != nil {
		log.WithFields(logrus.Fields{
			"operation": "Ambil User",
			"layer":     "Redis",
		}).WithError(err).Error("Data tidak ada")
		return Profil{}

	}

	// 	err := r.Rdb.HSet(ctx, key, map[string]any{
	// 	"id":       draft.ID,
	// 	"name":     draft.Name,
	// 	"username": draft.Username,
	// 	"email":    draft.Email,
	// 	"password": draft.Password,
	// 	"prodi":    draft.Prodi,
	// }).Err()

	id, _ := strconv.Atoi(userRes["id"])
	respone := Profil{
		ID:       id,
		Name:     userRes["name"],
		Username: userRes["username"],
		Password: userRes["password"],
		Prodi:    userRes["prodi"],
		Email:    userRes["email"],
	}
	return respone

}

func (r *implnRedisUser) GetOTP(ctx context.Context, key string) DataOTP {
	log := logger.New()

	if key == "" {

		log.WithFields(logrus.Fields{

			"operation": "Get OTP",
			"layer":     "Redis",
		}).Error("key tidak boleh kosong")
		return DataOTP{}
	}

	//Ambil string dari GET otp
	data, err := r.Rdb.Get(ctx, key).Result()
	log.WithFields(logrus.Fields{
		"operation": "GET OTP",
		"layer":     "Redis",
	}).Info(fmt.Sprintf("isi key: %s\n", key))
	if err != nil {

		log.WithFields(logrus.Fields{
			"operation": "Get OTP",
			"layer":     "redis",
		}).WithError(err).Error("data gagal di ambil")

		return DataOTP{}

	}

	tp := DataOTP{}

	err = json.Unmarshal([]byte(data), &tp)

	if err != nil {

		log.WithFields(logrus.Fields{})

		return DataOTP{}
	}

	return tp
}
