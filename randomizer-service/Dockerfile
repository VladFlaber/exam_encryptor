FROM  golang:alpine as builder
RUN mkdir /app
WORKDIR  /app
COPY . .
#ENV GO111MODULE=on
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/randomizer-service .
ENTRYPOINT ["./randomizer-service"]
