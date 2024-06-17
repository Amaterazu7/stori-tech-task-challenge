.PHONY: build clean #deploy

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/processor cmd/functions/processor_handler.go


clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	#sls deploy --verbose

format:
	gofmt -w cmd/functions/processor_handler.go

## S3
upload-csv:
	aws s3 cp ../source/transaction_list.csv s3://transaction-processor-bucket \
    	--endpoint-url http://localhost:4566
	aws --endpoint=http://localhost:4566/ s3 ls
list-s3:
	aws --endpoint-url http://localhost:4566 s3 ls s3://transaction-processor-bucket

## APP
start: clean build
	# AWS_ACCESS_KEY_ID=S3RVER AWS_SECRET_ACCESS_KEY=S3RVER sls offline start --host 0.0.0.0
	sls offline start --host 0.0.0.0
	# sls offline start --stage development --region us-east-1
