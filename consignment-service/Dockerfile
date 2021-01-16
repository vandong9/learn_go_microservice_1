FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD consignment-service /app/consignment-service

# COPY --from=builder /app/consignment-service .


CMD ["./consignment-service"]