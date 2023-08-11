FROM golang:1.3.3
COPY . /app
WORKDIR /app
RUN go build -o app
ENTRYPOINT [ "bash" ]
CMD [ "-c app" ]