# -------- Build stage --------
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary with a different name
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app-binary ./server

# -------- Runtime stage --------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/app-binary /app/app-binary
# Copy static assets into the runtime image so the server can serve index.html
COPY --from=builder /app/static /app/static

EXPOSE 3000

CMD ["/app/app-binary"]
