FROM golang:alpine

RUN mkdir -p /var/www

WORKDIR /var/www

COPY . /var/www

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go mod init github.com/snretuerma/gorest

RUN go mod tidy

RUN go build -o app

EXPOSE 3200

ENTRYPOINT [ "/var/www/app" ]