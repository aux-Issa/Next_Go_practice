FROM golang:1.18

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

ENV GD111MODULE=on

COPY . .
EXPOSE 1323
RUN go install github.com/cosmtrek/air@latest
CMD "air"