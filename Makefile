test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=cover.out

build_and_run:
	go build -o go-github .
	./go-github

swagger:
	go mod tidy
	go get -v github.com/swaggo/swag/cmd/swag
	go get -v github.com/swaggo/gin-swagger
	go get -v github.com/swaggo/files
	swag init -g main.go router/router.go --output docs