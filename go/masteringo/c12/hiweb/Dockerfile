FROM golang:latest

RUN mkdir /src
WORKDIR /src
COPY . .
RUN go build .
ENTRYPOINT [ "./hiweb" ]
