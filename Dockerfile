FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /ForumFinal

EXPOSE 8080

CMD ["/ForumFinal"]