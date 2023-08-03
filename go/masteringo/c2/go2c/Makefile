GOSRC := used_by_c.go
GOUT := used_by_c.o
CSRC := main.c
BIN := main

all: build_go build_c

build_go:
	go build -o ${GOUT} -buildmode=c-shared ${GOSRC}

build_c:
	gcc -o ${BIN} ${CSRC} ./${GOUT}
