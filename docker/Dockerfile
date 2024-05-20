FROM golang:1.22 AS builder

WORKDIR /build
RUN useradd -u 10001 gocanto

RUN go env -w GOPROXY="https://proxy.golang.org,direct"
RUN go env -w GONOSUMDB="github.com/gocanto/bridge*,github.com/gocanto/bridge/*/*"

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux
RUN go build -a -o main .

# Stage 2: Create a minimal production image
FROM golang:1.22 AS production
WORKDIR /app

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/main .
COPY --from=builder /build/config /app/config
COPY --from=builder /build/.env .

EXPOSE 8080
USER gocanto
CMD ["./main"]
