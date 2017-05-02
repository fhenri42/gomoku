NAME= algoGomokou
SRC =  srcs/server.go\
			 srcs/resolution.go\
			 srcs/expander.go\
			 srcs/helper.go\
			 srcs/graphicsInterphace.go
			#heuristique.go\

all: $(NAME)

$(NAME): $(SRC)
	go build -o  gomoku $(SRC)

clean:
	@echo 'No object files to delete'

fclean:
	rm -f $(NAME)

re: fclean all
