package main

import (
	"encoding/json"
	"fmt"
	"github.com/amaterazu7/transaction-processor/internal/application"
	"github.com/amaterazu7/transaction-processor/internal/infrastructure"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	const (
		SUCCESS = "SUCCESS"
		FAILED  = "CREDIT"
	)

	accountId := request.PathParameters["accountId"]
	log.Printf("[INFO] :: %s", accountId) // TODO:: REMOVE THIS ONLY to TESTS
	accountIdMsg := fmt.Sprintf("for AccountId { %s }", accountId)

	csvTransactionService := application.NewCsvTransactionService(
		infrastructure.NewPostgresAccountRepository(),
		infrastructure.NewPostgresTransactionRepository(),
	)
	statusCode, err := csvTransactionService.Run(accountId)

	if err != nil {
		message := fmt.Sprintf("API call failed %s: %s", accountIdMsg, err.Error())
		log.Printf("[ERROR] :: %s", message)
		marshalledErr, _ := json.Marshal(Response{
			Message: message,
			Status:  FAILED,
		})

		return events.APIGatewayProxyResponse{Body: string(marshalledErr), StatusCode: statusCode}, nil
	}

	marshalledResponse, _ := json.Marshal(Response{
		Message: fmt.Sprintf("Transaction processed %s", accountIdMsg),
		Status:  SUCCESS,
	})

	return events.APIGatewayProxyResponse{Body: string(marshalledResponse), StatusCode: statusCode}, nil
}

func main() {
	lambda.Start(Handler)
}
