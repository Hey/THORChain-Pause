package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Password string
	MakeCWD  string
}

var config Config

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	loadConfig()

	http.HandleFunc("/pause/", makeHandler("pause"))
	http.HandleFunc("/test/", makeHandler("status"))

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func loadConfig() {
	config = Config{
		Password: os.Getenv("PASSWORD"),
		MakeCWD:  os.Getenv("MAKE_CWD"),
	}
	if config.Password == "" || config.MakeCWD == "" {
		log.Fatal("PASSWORD and MAKE_CWD environment variables are required")
	}
	log.Printf("Config loaded: %+v\n", config)
}

func makeHandler(action string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSegments := strings.Split(r.URL.Path, "/")
		if len(pathSegments) < 3 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		password := pathSegments[2]

		if password != config.Password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		cmd := exec.Command("make", action)
		cmd.Dir = config.MakeCWD
		output, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, "Failed to execute command", http.StatusInternalServerError)
			log.Printf("Error executing make %s: %v", action, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(output)
	}
}
