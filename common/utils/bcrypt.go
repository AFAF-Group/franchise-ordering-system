package utils

import "golang.org/x/crypto/bcrypt"

const bcryptCost = 12

type BCrypt struct {
}

func NewBCrypt() *BCrypt {
	return &BCrypt{}
}

func (bc *BCrypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

func (bc *BCrypt) CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
