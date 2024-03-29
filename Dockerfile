FROM golang:1.16-alpine

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /app

EXPOSE 8080

CMD [ "./go-rest-api" ]
