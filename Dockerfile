FROM golang:alpine

WORKDIR /app/

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o backend_side .

EXPOSE 8080

CMD ["./backend_side"]