test:
	go test ./...

test-race:
	go test -race ./...

test-verbose:
	go test -v ./...

lint:
	golangci-lint run

coverage:
	go test -coverprofile=c.out ./... && go tool cover -html="c.out" && rm c.out

format:
	go fmt ./...
