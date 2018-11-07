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
	// ErrNotParsed is thrown when the json is not parsed
	ErrNotParsed = errors.New("error parsed json response")
)

//Response struct
type Response struct {
	Message string `json:"msg"`
}

//CustomRequest struct
type CustomRequest struct {
	Name string `json:"name"`
}

// Handler is your Lambda function handler
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		// return events.APIGatewayProxyResponse{
		// 	Body:       "ErrNameNotProvided",
		// 	StatusCode: 403}, ErrNameNotProvided

		res, err := createResponse("ErrNameNotProvided")

		if err != nil {

			return events.APIGatewayProxyResponse{}, err
		}

		return events.APIGatewayProxyResponse{Body: string(res),
			StatusCode: 403}, nil
	}

	var creq CustomRequest
	err := json.Unmarshal([]byte(request.Body), &creq)

	if err != nil {

		res, err := createResponse("ErrNotParsed")

		if err != nil {

			return events.APIGatewayProxyResponse{}, err
		}

		return events.APIGatewayProxyResponse{Body: string(res),
			StatusCode: 403}, nil

		// return events.APIGatewayProxyResponse{}, ErrNotParsed
	}

	//Making response
	m := Response{"Hello " + creq.Name}
	msg, err := json.Marshal(m)

	if err != nil {

		//return events.APIGatewayProxyResponse{Body: "ErrNotParsed",
		//	StatusCode: 403}, ErrNotParsed

		res, err := createResponse("ErrNotParsed")

		if err != nil {

			return events.APIGatewayProxyResponse{}, err
		}

		return events.APIGatewayProxyResponse{Body: string(res),
			StatusCode: 403}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(msg),
		StatusCode: http.StatusOK,
	}, nil

}

func createResponse(message string) ([]byte, error) {
	m := Response{message}
	msg, err := json.Marshal(m)

	if err != nil {
		return nil, ErrNotParsed
	}

	return msg, nil
}

func main() {
	lambda.Start(Handler)
}
