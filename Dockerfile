FROM golang:1.18-alpine3.15 AS builder

WORKDIR /app

COPY . /app

# configure go mod 
RUN go mod download

# build
RUN go build .

#multistage
FROM alpine:3.15

WORKDIR /

COPY --from=builder /app /app

CMD "pixelate"
