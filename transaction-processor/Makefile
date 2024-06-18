.PHONY: build clean #deploy

# =================================================================================================================== #
# Vars
# =================================================================================================================== #
FUNCTION_PATH 	:= cmd/functions
HANDLER_PATH 	:= processor_handler.go
FILE_NAME 		:= transaction_list.csv
ACCOUNT_ID 		:= 17d340fa-5bf5-4429-8167-bafe4c0af0a7
S3_BUCKET_NAME 	:= transaction-processor-bucket
S3_BUCKET_URL 	:= http://localhost:4566/

# =================================================================================================================== #
# AWS S3
# =================================================================================================================== #
upload-csv:
	aws s3 cp ../source/${FILE_NAME} s3://${S3_BUCKET_NAME}/${ACCOUNT_ID}_${FILE_NAME} \
    	--endpoint-url ${S3_BUCKET_URL}
list-s3:
	aws --endpoint=${S3_BUCKET_URL} s3 ls
	aws --endpoint-url ${S3_BUCKET_URL} s3 ls s3://${S3_BUCKET_NAME}

# =================================================================================================================== #
# Serverless
# =================================================================================================================== #
format:
	gofmt -w ${FUNCTION_PATH}/${HANDLER_PATH}

clean:
	rm -rf ./bin ./vendor Gopkg.lock

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/processor ${FUNCTION_PATH}/${HANDLER_PATH}

start-local: format clean build
	sls offline start --host 0.0.0.0
	# sls offline start --stage development --region us-east-1

deploy: clean build
	#sls deploy --verbose
