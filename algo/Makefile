NAME= algoGomokou
SRC =  server.go\
			 resolution.go\
			 findMove.go

all: $(NAME)

$(NAME): $(SRC)
	go build -o  algoGomokou $(SRC)

clean:
	@echo 'No object files to delete'

fclean:
	rm -f $(NAME)

re: fclean all
