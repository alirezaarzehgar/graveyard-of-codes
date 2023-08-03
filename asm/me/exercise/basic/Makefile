SRC := $(wildcard *.asm)
BINS := $(SRC:.asm=)

all: $(BINS)

%: %.asm
	nasm -f elf64 $< -o $@.o
	ld $@.o -o $@

clean:
	rm -rf $(BINS) *.o

.PHONE: all clean
