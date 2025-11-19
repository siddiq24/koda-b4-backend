package libs

import (
	"fmt"

	"github.com/matthewhartstonge/argon2"
)

func Generate_Hash(password string) string {
	fmt.Println(password)
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		fmt.Printf("failed hashing password: %x", err)
		return ""
	}

	return string(encoded)
}

func Verify_Hash(plainPassword, hashedPassword string) bool {
	fmt.Println(plainPassword, hashedPassword)
	if hashedPassword == "" {
		fmt.Printf("hashed password in database is empty")
		return false
	}

	ok, err := argon2.VerifyEncoded([]byte(plainPassword), []byte(hashedPassword))
	if err != nil {
		fmt.Printf("failed verifying password: %x", err)
		return false
	}

	return ok
}
