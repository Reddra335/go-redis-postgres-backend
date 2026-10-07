package pass

import "github.com/alexedwards/argon2id"

func GenerateHash(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

func ComparePassword(password, hash string) (bool, error) {

	return argon2id.ComparePasswordAndHash(password, hash)

}
