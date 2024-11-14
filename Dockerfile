FROM golang:1.23-alpine3.20

WORKDIR /rest_wallet

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o rest_wallet ./cmd

EXPOSE 8080

CMD [ "./rest_wallet" ]