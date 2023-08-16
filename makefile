
.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test ./... -v

.PHONY: test/cover
test/cover:
	go test -coverprofile=coverage.out ./...

.PHONY: cover
cover:
	go tool cover -html=coverage.out
