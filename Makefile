NAME= gomoku

SRC =  srcs/main.go\
			 srcs/algo_resolution.go\
			 srcs/algo_expander.go\
			 srcs/event.go\
			 srcs/algo_heuristique_helper.go\
			 srcs/algo_minmax.go\
			 srcs/display.go\
			 srcs/algo_helper.go\
			 srcs/algo_heuristique.go


all: $(NAME)

$(NAME): $(SRC)
	go build -o  $(NAME) $(SRC)

clean:
	@echo 'No object files to delete'

fclean:
	rm -f $(NAME)

re: fclean all
