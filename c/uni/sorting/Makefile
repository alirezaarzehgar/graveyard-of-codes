CC := cc
CFLAGS := -Werror -Wall -Wextra -g

SRC := $(wildcard *.c)
OBJ := $(SRC:.c=.o)
BIN := sorting

all: ${BIN}

${BIN}: ${OBJ}
	$(CC) ${OBJ} -o ${BIN}

clean:
	rm -rf ${OBJ} ${BIN}

test:
	bats .
