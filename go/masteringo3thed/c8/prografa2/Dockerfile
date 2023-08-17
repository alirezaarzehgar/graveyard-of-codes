FROM golang:latest
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o server
ENTRYPOINT [ "/bin/bash" ]
CMD [ "-c", "./server" ]