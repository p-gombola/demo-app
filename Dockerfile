FROM --platform=linux/arm64 golang:1.24-alpine3.21

COPY go.mod /go
COPY main.go /go

CMD [ "go", "run", "main.go" ]