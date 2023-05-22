FROM golang:1.16-alpine

WORKDIR /app

COPY . .
RUN go build -o reciept_processor .
EXPOSE 8080
CMD ["./reciept_processor"]