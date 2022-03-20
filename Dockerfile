FROM golang:1.16-alpine

## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## Add this go mod download command to pull in any dependencies
RUN go mod download
## RUN the following commands for Swagger
RUN go install github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/gin-swagger
RUN go install github.com/swaggo/files
RUN swag init -g main.go router/router.go --output docs
## we run go build to compile the binary
## executable of our Go program
RUN go build -o go-github . 
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/go-github"]