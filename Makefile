test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=cover.out

build_and_run:
	go build -o go-github .
	./go-github