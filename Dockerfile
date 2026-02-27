FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /api-server ./main.go


FROM alpine:latest

WORKDIR /

COPY --from=builder /api-server /api-server

EXPOSE 8080

CMD ["/api-server"]
