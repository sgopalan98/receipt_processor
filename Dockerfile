FROM golang:1.19-alpine

WORKDIR /app

COPY . .
RUN go build -o reciept_processor .
EXPOSE 8080
CMD ["./reciept_processor"]