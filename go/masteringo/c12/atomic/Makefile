all: build

build:
	go build .

test: clean build
	bats .

clean:
	rm -rf atomic