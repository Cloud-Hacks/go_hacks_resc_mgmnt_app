## Build

FROM golang:1.18-alpine

WORKDIR /res-mgmt

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /web4app

## Deploy

EXPOSE 8080

CMD [ "/web4app" ]
