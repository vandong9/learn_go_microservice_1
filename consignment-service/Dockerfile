FROM golang:alpine as builder
RUN apk --no-cache add git

WORKDIR /app/consignment-service

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

#######
FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
# ADD consignment-service /app/consignment-service

COPY --from=builder /app/consignment-service .


CMD ["./consignment-service"]