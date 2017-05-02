NAME= gomoku

SRC =  srcs/main.go\
			 srcs/algo_resolution.go\
			 srcs/algo_expander.go\
			 srcs/helper.go\
			 srcs/event.go\
			 srcs/helper_array.go\
			 srcs/helper_minmax.go\
			 srcs/helper_interface.go\
			 srcs/algo_heuristique.go

all: $(NAME)

$(NAME): $(SRC)
	go build -o  $(NAME) $(SRC)

clean:
	@echo 'No object files to delete'

fclean:
	rm -f $(NAME)

re: fclean all
