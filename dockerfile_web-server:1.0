FROM golang:1.24

WORKDIR /school_server

COPY go.mod go.sum /school_server/

RUN go mod download

COPY . ./

CMD ["go","run","main.go"]