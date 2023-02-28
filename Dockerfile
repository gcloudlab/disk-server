FROM golang:1.18 AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  GOPROXY=https://goproxy.cn

WORKDIR /app

COPY . .

RUN go mod download 
RUN go build -o app core/core.go

# Runer
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/app .
COPY core /app

# for Back4App 
EXPOSE 20088

CMD ["./app"]