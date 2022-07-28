FROM golang:1.18-alpine3.15 AS builder

WORKDIR /app

COPY . .

# configure go mod 
RUN go mod download

# build
RUN go build .

#multistage
FROM alpine:3.15

COPY --from=builder /app /app

CMD [ "/app/pixelate" ]
