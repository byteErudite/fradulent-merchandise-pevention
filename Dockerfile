FROM golang:1.16

WORKDIR /chaincode

COPY . .

RUN go mod vendor

CMD ["sh", "-c", "go run chaincode.go"]