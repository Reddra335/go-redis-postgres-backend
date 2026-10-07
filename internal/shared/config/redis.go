package config

import (
	"context"
	"encoding/json"
	"errors"
	profilmahasiswa "newProjek/internal/profil_mahasiswa"
	domainredis "newProjek/internal/shared/config/domain_redis"
	"newProjek/internal/shared/logger"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisUser interface {
	CreateUser(ctx context.Context, do profilmahasiswa.Profil, key string) string
	GetUserEmail(ctx context.Context, key string) string
	SimpanOTP(ctx context.Context, email string, otp string, keyUser string) map[string]any // menambahkan otp sekaligus mengembalikan otp nya, redis string
	GetUser(ctx context.Context, key string) profilmahasiswa.Profil
}

type ImplnRedisUser struct {
	Rdb *redis.Client
}

func NewImplnRedisUser(rdb *redis.Client) RedisUser {

	return &ImplnRedisUser{Rdb: rdb}
}

//SELECT id, name, username, email, password, prodi FROM mahasiswa.profil WHERE id=2;

// Create User
func (r *ImplnRedisUser) CreateUser(ctx context.Context, do profilmahasiswa.Profil, key string) string {

	log := logger.New()
	if key == "" {

		err := errors.New("Key Anda bermassalah")
		log.WithFields(logrus.Fields{
			"operation": "Create User Redis",
			"layer":     "Redis",
		}).Fatal(err)
	}
	draft := domainredis.UserDraft{

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
		}).WithError(err).Error("terjadi kesalahan dalam create redis")
		return ""
	}

	errX := r.Rdb.Expire(ctx, "uji", 30*time.Minute).Err()

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

func (r *ImplnRedisUser) GetUserEmail(ctx context.Context, key string) string {
	log := logger.New()
	if key == "" {

		err := errors.New("Key Anda bermasalah")
		log.WithFields(logrus.Fields{
			"operation": "Get Email",
			"layer":     "Redis",
		}).Fatal(err)
		return ""
	}
	email, err := r.Rdb.HGet(ctx, key, "email").Result()
	if err != nil {
		log.WithFields(logrus.Fields{
			"operation": "Get Email",
			"layer":     "Redis",
		}).WithError(err).Fatal("get email redis bermasalah")
		return ""
	}
	log.WithFields(logrus.Fields{
		"operation": "Get Email",
		"layer":     "Redis",
	}).Info("Data berhasil diambil")
	return email
}

// Simpan OTP dan ambil
func (r *ImplnRedisUser) SimpanOTP(ctx context.Context, email string, otp string, keyUser string) map[string]any {
	log := logger.New()

	if email == "" {

		log.WithFields(logrus.Fields{
			"operation": "create otp",
			"layer":     "Redis",
		}).Fatal("Kosong")
		return nil
	}

	if otp == "" {

		log.WithFields(logrus.Fields{
			"operation": "create otp",
			"layer":     "Redis",
		}).Fatal("Kosong")
		return nil
	}

	data := map[string]any{

		"user": keyUser,
		"otp":  otp,
	}

	dataJ, errJ := json.Marshal(data)

	if errJ != nil {
		log.WithFields(logrus.Fields{
			"operation": "json redis",
			"layer":     "redis",
		}).WithError(errJ).Error("gagal mengubah ke json")
		return nil

	}

	ob := string(dataJ)

	//Simpan OTP
	err := r.Rdb.Set(ctx, keyUser, ob, 5*time.Minute).Err()
	if err != nil {
		log.WithFields(logrus.Fields{

			"operation": "create otp",
			"layer":     "Redis",
		}).WithError(err).Fatal("ada yang salah dengan create otp")
		return nil
	}

	log.Info("OTP berhasil di simpan")

	res, err := r.Rdb.Get(ctx, keyUser).Result()
	if err != nil {
		log.WithFields(logrus.Fields{
			"operation": "Ambil data OTP",
			"layer":     "Redis",
		}).WithError(err).Error("Gagal mengambil OTP")
		return nil
	}
	//Ambil OTP
	var ambilUser map[string]any

	json.Unmarshal([]byte(res), &ambilUser)
	return ambilUser
}

// Ambil User
func (s *ImplnRedisUser) GetUser(ctx context.Context, key string) profilmahasiswa.Profil {
	log := logger.New()
	if key == "" {

		errK := errors.New("key kosong")
		log.WithFields(logrus.Fields{
			"operation": "Ambil User",
			"key":       "Redis",
		}).WithError(errK)

		return profilmahasiswa.Profil{}
	}

	userRes, err := s.Rdb.HGetAll(ctx, key).Result()

	if err != nil {
		log.WithFields(logrus.Fields{
			"operation": "Ambil User",
			"layer":     "Redis",
		}).WithError(err).Error("Data tidak ada")
		return profilmahasiswa.Profil{}

	}

	userJSON, errJ := json.Marshal(userRes)
	if errJ != nil {
		log.WithFields(logrus.Fields{
			"operation": "Ambil User",
			"layer":     "Redis",
		}).WithError(errJ).Error("gagal mengubah data")
	}
	res := profilmahasiswa.Profil{}

	errJSON := json.Unmarshal(userJSON, &res)

	if errJSON != nil {
		log.WithFields(logrus.Fields{
			"operation": "Ambil User",
			"layer":     "Redis",
		}).WithError(errJSON).Error("Gagal mengubah data json")
		return profilmahasiswa.Profil{}

	}
	return res

}
