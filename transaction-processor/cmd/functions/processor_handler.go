package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/amaterazu7/transaction-processor/internal/application"
	"github.com/amaterazu7/transaction-processor/internal/infrastructure"
	"github.com/amaterazu7/transaction-processor/internal/infrastructure/persistence"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	const (
		SUCCESS = "SUCCESS"
		FAILED  = "FAILED"
	)

	log.Printf("[INFO] :: Starting Lambda Handler :: ")
	accountId := request.PathParameters["accountId"]
	accountIdMsg := fmt.Sprintf("for AccountId { %s }", accountId)
	err := godotenv.Load()
	if err != nil {
		log.Printf("[ERROR] :: Loading .env file failed :: ")
	}

	dbConn, dbConnErr := persistence.NewDBConfig().ConnectToDB()
	if dbConnErr != nil {
		log.Printf("[ERROR] :: %s", dbConnErr.Error())
		marshalledErr, _ := json.Marshal(Response{
			Message: dbConnErr.Error(),
			Status:  FAILED,
		})

		return events.APIGatewayProxyResponse{Body: string(marshalledErr), StatusCode: 500}, nil
	}
	defer dbConn.Close()

	csvTransactionService := application.NewCsvTransactionService(
		accountId,
		infrastructure.NewMySqlAccountRepository(dbConn),
		infrastructure.NewMySqlTransactionRepository(dbConn, context.Background()),
		persistence.NewS3BucketRepository(os.Getenv("AWS_S3_BUCKET_NAME"), os.Getenv("AWS_REGION")),
	)
	statusCode, processorResult, processorErr := csvTransactionService.RunProcessor()

	statusCode, senderErr := application.NewEmailSenderService().SendMessage(processorResult)

	if processorErr != nil || senderErr != nil {
		var msg string
		switch {
		case processorErr != nil:
			msg = fmt.Sprintf("API processor failed %s: %s", accountIdMsg, processorErr.Error())
		default:
			msg = fmt.Sprintf("API Sender failed %s: %s", accountIdMsg, senderErr.Error())
		}

		log.Printf("[ERROR] :: %s", msg)
		marshalledErr, _ := json.Marshal(Response{
			Message: msg,
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
