package main

import (
	"os"

	"github.com/ChqThomas/gitlab-webooks-awtrix3/api"
	"github.com/ChqThomas/gitlab-webooks-awtrix3/webhook"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load(".env")

	client := api.NewClient(
		os.Getenv("AWTRIX_BASE_URL"),
		os.Getenv("AWTRIX_USERNAME"),
	    os.Getenv("AWTRIX_PASSWORD"),
	)

	client.SendNotification("Starting Gitlab webhooks server", "#FFFFFF", false)

	webhook.ListenWebhook(client)
}