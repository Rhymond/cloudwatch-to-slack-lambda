build:
	GOARCH=amd64 GOOS=linux go build -o ./bin/cmd -v ./cmd/main.go

deploy: build
	sam deploy
