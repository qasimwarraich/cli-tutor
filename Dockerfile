FROM golang:1.18

WORKDIR /cli-tutor

COPY go.mod go.sum ./

RUN go mod download && go mod verify

ENV TERM xterm-256color

ADD ./ ./

RUN go build -v -o ./cli-tutor ./src/cli-tutor/...

CMD ./cli-tutor
