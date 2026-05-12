FROM golang:1.25.6-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
# proxy.golang.org kadang putus di CI; direct sebagai fallback + beberapa retry
ENV GOPROXY=https://proxy.golang.org,direct
RUN ok=0; for i in 1 2 3 4 5; do \
  go mod download && ok=1 && break; \
  echo "go mod download failed, retry $$i/5 in 15s..."; \
  sleep 15; \
done; test "$$ok" = 1
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server ./cmd/main.go

FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8081
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8081/health || exit 1
CMD ["./server"]