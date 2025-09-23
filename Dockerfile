FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-server ./main.go


FROM alpine:latest

WORKDIR /

COPY --from=builder /api-server /api-server

EXPOSE 8080

CMD ["/api-server"]
