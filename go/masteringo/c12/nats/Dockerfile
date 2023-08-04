FROM golang:latest
RUN mkdir /app
COPY ./src /app
WORKDIR /app
RUN go build .
ENTRYPOINT [ "./hinats" ]
