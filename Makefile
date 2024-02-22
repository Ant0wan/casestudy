GO   := go
NAME := myprogram

all: build

build:
	CGO_ENABLED=0 GOOS=linux $(GO) build -v -o $(NAME)

test:
	sh test.sh
clean:
	rm -rf $(NAME)

fclean: clean

.PHONY: all build  test clean fclean
