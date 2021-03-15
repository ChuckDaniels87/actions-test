FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN go build -v

FROM alpine
WORKDIR /app
COPY --from=builder /app/first /bin/first

ENTRYPOINT ["first"]
