FROM golang:1.24

WORKDIR /app

RUN go mod init test

COPY . .

RUN go build -o math

CMD [ "./math" ]