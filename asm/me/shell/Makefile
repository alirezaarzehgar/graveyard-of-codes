SRC := $(wildcard *.asm)
OBJ := $(SRC:.asm=.o)
TARGET := shell

all: $(TARGET)

$(TARGET): $(OBJ)
	ld $< -o $@

$(OBJ): $(SRC)
	nasm -f elf64 $< -o $@

clean:
	rm -rf $(TARGET) $(OBJ)
