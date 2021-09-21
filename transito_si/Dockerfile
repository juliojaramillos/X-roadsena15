FROM golang:latest

WORKDIR /go/src/unal.edu.co/rest-example/
RUN go version

COPY ./transito_si /go/src/unal.edu.co/rest-example/

EXPOSE 1000
EXPOSE 8080

CMD ["make","dev"]
