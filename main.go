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
	// ErrNotParsed is thrown when the json is not parsed
	ErrNotParsed = errors.New("error parsed json")
	// ErrNotBody is thrown when a body is not provided
	ErrNotBody = errors.New("nothing was provided in the HTTP body")
)

//JSON type
type JSON []byte

//APIRequest type
type APIRequest events.APIGatewayProxyRequest

//APIResponse type
type APIResponse events.APIGatewayProxyResponse

//Response struct
type Response struct {
	Message string `json:"msg"`
}

//Input struct
type Input struct {
	Name string `json:"name"`
}

// Handler is your Lambda function handler
func Handler(request APIRequest) (APIResponse, error) {

	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {

		return throwAPIError(ErrNotBody)
	}

	//Decoding JSON
	var in Input
	err := json.Unmarshal(JSON(request.Body), &in)

	if err != nil {

		return throwAPIError(err)
	}

	//Coding JSON
	msg, err := createResponseJSON("Hello " + in.Name)

	if err != nil {
		return throwAPIError(err)
	}

	return APIResponse{
		Body:       string(msg),
		StatusCode: http.StatusOK,
	}, nil

}

//Aux func to throw errors
func throwAPIError(err error) (APIResponse, error) {

	res, err := createResponseJSON(err.Error())

	if err != nil {

		return throwAPIError(err)
	}

	return APIResponse{Body: string(res),
		StatusCode: 403}, nil
}

//Aux func to create a ResponseJSON
func createResponseJSON(message string) (JSON, error) {
	m := Response{message}
	msg, err := json.Marshal(m)

	return msg, err
}

func main() {
	lambda.Start(Handler)
}
