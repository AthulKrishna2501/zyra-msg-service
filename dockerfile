FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .  

RUN go build -o msg-service ./cmd  

EXPOSE 5003

CMD ["./msg-service"]
