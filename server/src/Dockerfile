FROM golang:1.16-alpine

WORKDIR /go/src/github.com/39TO/gockerql

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]