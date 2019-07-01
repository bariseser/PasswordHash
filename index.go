package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Make(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Validate(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "your password"
	hash, _ := Make(password)

	fmt.Println("Password:", password)
	fmt.Println("hash:", hash)

	match := Validate(password, hash)
	fmt.Println("Match:", match)
}