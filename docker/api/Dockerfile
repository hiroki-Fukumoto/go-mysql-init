FROM golang:1.16.2-buster

ENV GO111MODULE=auto

WORKDIR /go/api

COPY ./api .

RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/gin-gonic/gin && \
  go get -u github.com/go-delve/delve/cmd/dlv

COPY ./api/go.mod ./api/go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
