FROM golang:1.21.3
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o app
ENTRYPOINT [ "bash", "-c" ]
CMD [ "./app" ]
