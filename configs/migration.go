package configs

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func MigrationsUp() {
	env := GetEnv()
	cmd := exec.Command(
		"migrate",
		"-path", "migrations/up",
		"-database",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			env.POSTGRES_USER,
			env.POSTGRES_PASSWORD,
			env.POSTGRES_HOST,
			env.POSTGRES_PORT,
			env.POSTGRES_NAME,
		),
		"up",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Gagal menjalankan migrasi: %v", err)
	}

	log.Println("Migrasi database selesai!")
}
