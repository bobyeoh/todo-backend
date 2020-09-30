FROM golang:latest
LABEL maintainer="Bob Yang<b@001.gs>"
WORKDIR $GOPATH/src/todo-backend
COPY . $GOPATH/src/todo-backend
RUN go build main.go
EXPOSE 8888
ENTRYPOINT ["./main"]
