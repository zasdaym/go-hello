FROM golang:1.14 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
ENV PORT 8080
EXPOSE PORT
CMD ["./app"]
