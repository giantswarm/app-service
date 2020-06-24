FROM alpine:3.12

RUN apk add --no-cache ca-certificates

ADD ./app-service /app-service

ENTRYPOINT ["/app-service"]
