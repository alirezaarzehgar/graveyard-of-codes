SRC := hello.asm
OBJ := $(SRC:.asm=.o)
TARGET := hello

all: $(TARGET)

$(TARGET): $(OBJ)
	ld $(OBJ) -o $(TARGET)

$(OBJ): $(SRC)
	nasm -f elf64 $< -o $@

clean:
	rm -rf $(OBJ) $(TARGET)
