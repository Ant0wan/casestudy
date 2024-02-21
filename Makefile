GO   := go
NAME := myprogram

all: build

build:
	$(GO) build -o $(NAME)

test:
	sh test.sh
clean:
	rm -rf $(NAME)

fclean: clean

.PHONY: all build  test clean fclean
