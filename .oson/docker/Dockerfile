FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o theoson .

FROM alpine:latest

RUN apk --no-cache add ca-certificates


WORKDIR /app

COPY --from=builder app/theoson .
EXPOSE 8080

CMD ["./theoson"]

