package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var TELEGRAM_BOT_TOKEN string

type Data struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "Hello There...",
		}, nil
	}

	var err error

	data := Data{}

	if err = json.Unmarshal([]byte(req.Body), &data); err != nil {

		log.Println(err)

		return events.APIGatewayProxyResponse{
			StatusCode: 503,
		}, nil

	}

	if strings.Contains(data.Message.Text, "start") {
		return events.APIGatewayProxyResponse{
			StatusCode: 204,
		}, nil
	}

	responseData := map[string]interface{}{
		"text":    data.Message.Text,
		"chat_id": data.Message.Chat.ID,
	}

	var responseDataJSON []byte

	if responseDataJSON, err = json.Marshal(responseData); err != nil {

		log.Println(err)

		return events.APIGatewayProxyResponse{
			StatusCode: 503,
		}, nil

	}

	if request, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", TELEGRAM_BOT_TOKEN), bytes.NewReader(responseDataJSON)); err != nil {

		log.Println(err)

		return events.APIGatewayProxyResponse{
			StatusCode: 503,
		}, nil

	} else {

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		if _, err = client.Do(request); err != nil {

			log.Println(err)

			return events.APIGatewayProxyResponse{
				StatusCode: 503,
			}, nil

		}

	}

	return events.APIGatewayProxyResponse{
		StatusCode: 204,
	}, nil

}

func main() {
	lambda.Start(handler)
}
