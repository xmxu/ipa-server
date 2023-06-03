# builder
FROM golang:1.16 AS builder
ENV GOPROXY=https://proxy.golang.com.cn,direct
ENV GOARCH=amd64
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -ldflags '-linkmode "external" --extldflags "-static"' cmd/ipasd/ipasd.go

# runtime
FROM ineva/alpine:3.10.3
WORKDIR /app
COPY --from=builder /src/ipasd /app
COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT /docker-entrypoint.sh
