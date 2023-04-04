FROM golang:alpine3.17

WORKDIR /app

ENV GIN_MODE=release

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . ./

#RUN wget -O - https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem > ./rds-combined-ca-bundle.pem

RUN go build -o /cs440-keylogger/go-gin-backend

EXPOSE 8080

CMD [ "/cs440-keylogger/go-gin-backend" ]
