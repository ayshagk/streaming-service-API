package utils

import "golang.org/x/crypto/bcrypt"

//hash and compare passwords, so we can store password in DB safely

func HashPassword(password string) (string, error) {
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(HashPassword), nil 
}

func ComparePassword(plainPass, HashedPass string) error{
	err := bcrypt.CompareHashAndPassword([]byte(HashedPass), []byte(plainPass))
	if err != nil {
		return err
	}
	
	return nil
}