NAME= gomoku

SRC =  srcs/main.go\
			 srcs/resolution.go\
			 srcs/expander.go\
			 srcs/helper.go\
			 srcs/initInterphace.go\
			 srcs/display.go
			#heuristique.go\

all: $(NAME)

$(NAME): $(SRC)
	go build -o  $(NAME) $(SRC)

clean:
	@echo 'No object files to delete'

fclean:
	rm -f $(NAME)

re: fclean all
