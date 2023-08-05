all:

build:
	go build .
	docker-compose build

up: down build
	docker-compose up --force-recreate

down: clean
	docker-compose down

clean:
	rm -rf hi