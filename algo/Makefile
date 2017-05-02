NAME= algoGomokou
SRC =  main.go\
			 resolution.go\
			 expander.go\
			 helper.go\
			 initInterphace.go\
			 display.go
			#heuristique.go\

all: $(NAME)

$(NAME): $(SRC)
	go build -o  algoGomokou $(SRC)

clean:
	@echo 'No object files to delete'

fclean:
	rm -f $(NAME)

re: fclean all
