FROM golang:1.18-alpine3.15 AS builder

WORKDIR /pixelate-app

COPY . /pixelate-app

# configure go mod 
RUN go mod download

# build
RUN go build -o /pixelate-bot

#multistage
FROM alpine:3.15
COPY --from=builder /pixelate-app /pixelate-app

CMD ["./pixelate-bot"]
