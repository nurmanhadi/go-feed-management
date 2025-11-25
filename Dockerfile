FROM docker.io/golang:1.25.4-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./main.go

FROM docker.io/library/debian:stable-slim
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 3002
CMD [ "./app" ]