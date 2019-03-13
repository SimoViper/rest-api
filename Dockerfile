FROM golang

WORKDIR /go/src

COPY ./ /go/src/rest-api/

RUN go get -d github.com/gorilla/mux

RUN cd /go/src/rest-api/src && go build -o main

EXPOSE 8000

ENTRYPOINT "./rest-api/src/main"
