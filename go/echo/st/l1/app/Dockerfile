FROM golang:1.20.6
COPY . /app
WORKDIR /app
RUN go build -o app
ENTRYPOINT [ "./app" ]