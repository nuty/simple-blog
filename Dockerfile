FROM golang:1.21.1-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

EXPOSE 3001

CMD ["go", "run", "app.go"]
