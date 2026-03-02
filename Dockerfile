FROM golang

WORKDIR /app

COPY . .

ENV GO111MODULE=off

RUN go build -o app

CMD ["./app"]

