package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
	// ErrResponseNotParsed is thrown when the response json is not parsed
	ErrResponseNotParsed = errors.New("error parsed json response")
)

//Response struct
type Response struct {
	Message string `json:"msg"`
}

// Handler is your Lambda function handler
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	m := Response{"Hello " + request.Body}
	msg, err := json.Marshal(m)

	if err != nil {

		return events.APIGatewayProxyResponse{}, ErrResponseNotParsed
	}

	return events.APIGatewayProxyResponse{
		Body:       string(msg),
		StatusCode: http.StatusOK,
	}, nil

}

func main() {
	lambda.Start(Handler)
}
