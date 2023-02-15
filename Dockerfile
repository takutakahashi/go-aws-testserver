FROM golang:1.20 as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app main.go

FROM ubuntu
WORKDIR /
COPY --from=builder /workspace/app .
RUN apt update && apt install -y ca-certificates

ENTRYPOINT ["/app"]
