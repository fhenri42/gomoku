NAME= gomoku

SRC =  srcs/main.go\
			 srcs/resolution.go\
			 srcs/expander.go\
			 srcs/helper.go\
			 srcs/interface.go\
			 srcs/display.go\
			 srcs/array_helper.go\
			 srcs/minmax_helper.go\
			 srcs/heuristique.go\

all: $(NAME)

$(NAME): $(SRC)
	go build -o  $(NAME) $(SRC)

clean:
	@echo 'No object files to delete'

fclean:
	rm -f $(NAME)

re: fclean all
