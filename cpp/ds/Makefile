CXX := g++
CFLAGS := -Werror -Wall -Wextra -g

SRC := $(wildcard *.cpp)
BIN := $(SRC:.cpp=)

all: ${BIN}

${BIN}:
	$(CXX) $@.cpp -o $@

clean:
	rm -rf ${BIN}
