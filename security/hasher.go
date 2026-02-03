package security

import (
	"github.com/alexedwards/argon2id"
)

func ComparePassword(password, hash string) (bool, error) {
	ok, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return ok, nil
}
