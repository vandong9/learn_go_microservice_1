FROM golang:alpine as builder

WORKDIR /go/src/github.com/vandong9/learn_go_microservice_1/user-cli
COPY . .

# RUN go get -u github.com/golang/dep/cmd/dep
# RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/vandong9/learn_go_microservice_1/user-cli/user-cli .
ENTRYPOINT ["./user-cli"]

CMD ["./user-cli"]