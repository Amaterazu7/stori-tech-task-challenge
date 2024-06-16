.PHONY: build clean #deploy

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/processor cmd/functions/processor_handler.go


clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	#sls deploy --verbose

format:
	gofmt -w cmd/functions/processor_handler.go

start: clean build
	sls offline start --host 0.0.0.0
	# sls offline start --stage development --region us-east-1
