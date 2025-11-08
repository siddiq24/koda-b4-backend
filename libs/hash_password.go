package libs

import (
	"log"

	"github.com/matthewhartstonge/argon2"
)

func Create_Hash(pass string) string {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(pass))
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(encoded)
}

func Verify_Hash(usePass, dbPass string) bool {
	log.Println(usePass, "\n", dbPass)
	ok, err := argon2.VerifyEncoded([]byte(usePass), []byte(dbPass))
	if err != nil {
		log.Println(err)
		return false
	}

	return ok
}
