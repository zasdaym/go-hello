FROM golang:1.13 as builder
WORKDIR /app
COPY . .
RUN go build -o app .

FROM alpine:3.9.4
WORKDIR /app
COPY --from=builder /app .
CMD ["./app"]
