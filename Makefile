NAME= gomoku

SRC =  srcs/main.go\
			 srcs/algo_resolution.go\
			 srcs/event.go\
			 srcs/algo_minmax.go\
			 srcs/display.go\
			 srcs/helper.go\
			 srcs/algo_check.go\
			 srcs/algo_simulateMove.go\
			 srcs/algo_heuristic.go\


all: $(NAME)

$(NAME): $(SRC)
	go build -o  $(NAME) $(SRC)

clean:
	@echo 'No object files to delete'

fclean:
	rm -f $(NAME)

re: fclean all
