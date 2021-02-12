FROM golang:alpine as builder
RUN apk --no-cache add git

WORKDIR /app/user-service

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service -a -installsuffix cgo main.go repository.go token_service.go handler.go database.go

#######
FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
# ADD user-service /app/user-service

COPY --from=builder /app/user-service .


CMD ["./user-service"]