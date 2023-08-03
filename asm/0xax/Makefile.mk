SRC := $(wildcard *.asm)
TARGETS := $(SRC:.asm=)

all: ${TARGETS}

%: %.asm
	nasm -f elf64 $< -o $@.o
	ld $@.o -o $@

clean:
	rm -rf ${TARGETS} *.o

.PHONY: all clean
