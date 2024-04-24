FROM golang:1.22.1 AS builder

WORKDIR /usr/chat-auth-api/app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install gotest.tools/gotestsum@latest
RUN apt-get install -y git curl

EXPOSE 3000

CMD [ "tail", "-f", "/dev/null" ]
