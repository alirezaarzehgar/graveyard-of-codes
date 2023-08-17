FROM golang:latest
RUN mkdir /app
COPY *.* /app
WORKDIR /app
RUN go build -o server
ENTRYPOINT [ "/bin/bash" ]
CMD [ "-c", "./server" ]
