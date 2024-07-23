FROM golang:1.19-alpine

# Install gcc and other necessary build tools
RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
