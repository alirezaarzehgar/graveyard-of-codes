BIN := hi

all: clean build up

up: down
	docker-compose up --force-recreate

down:
	docker-compose down

build: app-build docker-build

app-build:
	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o bin/

docker-build:
	docker-compose build

clean:
	rm -rf bin

clean-volums:
