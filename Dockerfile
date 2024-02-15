# Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY main.go favicon.png go.mod index.html .
RUN CGO_ENABLED=0 go build -o server main.go

# Final stage
FROM scratch
COPY --from=builder /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]