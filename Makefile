.PHONY:
.SILENT:

# building binary file
build:
	go build -o ./.bin/service ./cmd/main.go

run: build
	./.bin/service

build-image:
	docker build -t service_authorization:v1 .

start-container:
	docker run --name service-auth -p 82:82 --env-file .env service_authorization:v1