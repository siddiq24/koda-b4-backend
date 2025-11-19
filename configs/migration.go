package configs

import (
	"log"
	"os"
	"os/exec"
)

func MigrationsUp() {
	cmd := exec.Command(
		"migrate",
		"-path", "migrations/up",
		"-database",
		os.Getenv("PSQL_URL"),
		"up",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Gagal menjalankan migrasi: %v", err)
	}

	log.Println("Migrasi database selesai!")
}
