package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func loadEnv() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("❌ Failed to get executable path:", err)
	}
	exeDir := filepath.Dir(exePath)

	env := os.Getenv("APP_ENV")
	envFile := ".env"
	if env != "" {
		envFile = ".env." + env
	}

	envPath := filepath.Join(exeDir, envFile)
	if err := godotenv.Load(envPath); err != nil {
		if cwd, _ := os.Getwd(); cwd != "" {
			envPath = filepath.Join(cwd, envFile)
			if err := godotenv.Load(envPath); err != nil {
				log.Fatalf("❌ Error loading %s file", envPath)
			}
		} else {
			log.Fatal(errors.New("❌ cannot resolve working directory"))
		}
	}
}
