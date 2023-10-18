dep:
	go mod tidy
	go mod vendor

dev:
	go run .\main.go

build:
	set GOOS=linux && set GOARCH=amd64 && go build -o bin/talkjs .\main.go

client-dev:
	pnpm  run --prefix ./vue-talkjs dev