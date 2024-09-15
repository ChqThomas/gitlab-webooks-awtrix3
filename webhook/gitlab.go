package webhook

import (
	"fmt"
	"os"
	"strings"

	"net/http"

	"github.com/ChqThomas/gitlab-webooks-awtrix3/api"
	"github.com/go-playground/webhooks/v6/gitlab"
)

const (
	path = "/webhooks"
)

func ListenWebhook(client *api.Client) {
	hook, _ := gitlab.New(gitlab.Options.Secret(os.Getenv("GITLAB_WEBHOOK_SECRET")))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got a request")
		payload, err := hook.Parse(r, gitlab.PipelineEvents)
		fmt.Printf("%+v", payload)
		if err != nil {
			fmt.Printf("%+v", err)
			if err == gitlab.ErrEventNotFound {
				fmt.Println("Event not found")
			}
		}
		switch payload.(type) {

		case gitlab.PipelineEventPayload:
			pipeline := payload.(gitlab.PipelineEventPayload)
			color := "#FFFF00"
			hold := true
			message := pipeline.ObjectAttributes.Status
			if pipeline.ObjectAttributes.Status == "failed" {
				color = "#FF0000"
				hold = false
			} else if pipeline.ObjectAttributes.Status == "success" {
				color = "#00FF00"
				hold = false
			} else if pipeline.ObjectAttributes.Status == "running" {
				color = "#0000FF"
				message += " - " + strings.Join(pipeline.ObjectAttributes.Stages, ", ")
			}
			client.SendNotification(message, color, hold)
		}
	})

	fmt.Println("Listening on port " + os.Getenv("SERVER_PORT"))
	http.ListenAndServe(":" + os.Getenv("SERVER_PORT"), nil)
}