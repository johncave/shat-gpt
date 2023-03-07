FROM docker.io/golang:1.20.1

WORKDIR /elephant/

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o ./elephant

EXPOSE 8080

CMD ["/elephant/elephant"]