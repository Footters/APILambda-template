# Overview

API Gateway events consist of a request that was routed to a Lambda function by API Gateway. When this happens, API Gateway expects the result of the function to be the response that API Gateway should respond with.

## Preparing and deploying Function to Lambda
```
pip install awscli
go get
go test
cd cmd/
./build-go
aws lambda update-function-code --function-name example --zip-file fileb://main.zip
```
