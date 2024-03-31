NAME := gomoku

SRC := $(wildcard *.go)

GOBUILD := go build

TARGET := $(NAME)

.PHONY: all clean fclean re

all: $(TARGET)

$(TARGET): $(SRC)
	$(GOBUILD) -o $@

clean:
	rm -f $(TARGET)

fclean: clean

re: fclean all
