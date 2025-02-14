FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache ca-certificates busybox-extras

COPY --from=builder /app/main .

COPY init-app.sh .

RUN chmod +x ./main ./init-app.sh

EXPOSE 9090

CMD ["./init-app.sh"]
