package main

import (
	"flag"
	"os"

	"github.com/ChqThomas/gitlab-webooks-awtrix3/api"
	"github.com/ChqThomas/gitlab-webooks-awtrix3/webhook"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load(".env")

	initFlags()

	client := api.NewClient(
		os.Getenv("AWTRIX_BASE_URL"),
		os.Getenv("AWTRIX_USERNAME"),
	    os.Getenv("AWTRIX_PASSWORD"),
	)

	client.SendNotification("Starting Gitlab webhooks server", "#FFFFFF", false)

	webhook.ListenWebhook(client)
}

func initFlags() {
	baseUrl := flag.String("awtrix-host", "", "Base URL of the AWTRIX 3 server")
    username := flag.String("awtrix-username", "", "Username for the AWTRIX 3 server")
    password := flag.String("awtrix-password", "", "Password for the AWTRIX 3 server")
	gitlabSecret := flag.String("gitlab-secret", "", "Gitlab webhook secret")
	serverPort := flag.String("port", "", "Server port")

    // Parse les flags
    flag.Parse()

    // Si les flags ne sont pas fournis, on vérifie les variables d'environnement
    if *username != "" {
        os.Setenv("AWTRIX_USERNAME", *username)
    }
    if *password != "" {
        os.Setenv("AWTRIX_PASSWORD", *password)
    }
	if *baseUrl != "" {
		os.Setenv("AWTRIX_BASE_URL", *baseUrl)
	}
	if *gitlabSecret != "" {
		os.Setenv("GITLAB_WEBHOOK_SECRET", *gitlabSecret)
	}
	if *serverPort != "" {
		os.Setenv("SERVER_PORT", *serverPort)
	}
}