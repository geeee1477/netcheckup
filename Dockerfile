FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o netcheckup ./cmd/netcheckup

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/netcheckup .

ENTRYPOINT ["./netcheckup"]
