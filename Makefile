
generate:
	go run ./cmd/generate_testdata/main.go

test:
	go test ./...

benchmark:
	go test -bench=. ./...