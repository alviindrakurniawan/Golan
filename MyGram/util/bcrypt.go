package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) ([]byte, error) {
	
	return bcrypt.GenerateFromPassword(password, 8)
}

func HashMacthesPassword(hash []byte, plain []byte) bool {
	err:= bcrypt.CompareHashAndPassword(hash, plain)
	return err == nil
}


