package main

import (
	"fmt"
	"os"
	"strconv"
)

var black string = "#"
var white string = " "
var result string = ""

func main() {
	args := os.Args

	x, y := validate_args(args)

	fmt.Println("Программа для рисования шахматной доски указанного пользователем размера.")

	fmt.Println("Ширина - ", args[1])

	fmt.Println("Высота - ", args[2], "\n")

	for i := 1; i <= x; i++ {
		if i%2 == 0 {
			for j := 1; j <= y; j++ {
				if j%2 != 1 {
					result = result + black
				} else {
					result = result + white
				}
			}
		} else {
			for j := 1; j <= y; j++ {
				if j%2 != 1 {
					result = result + white
				} else {
					result = result + black
				}
			}
		}
		result = result + "\n"
	}

	fmt.Print(result)
}

func validate_args(args []string) (int, int) {

	if len(args) != 3 {
		fmt.Println("Неверное количество аргументов.")
		os.Exit(1)
	}

	x, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Println("Ошибка в аргументе, должно быть целое число.", err)
		os.Exit(1)
	}

	y, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Println("Ошибка в аргументе, должно быть целое число.", err)
		os.Exit(1)
	}

	if x == 0 || y == 0 {
		fmt.Println("Ширина или высота игрового поля не может быть равна 0.")
		os.Exit(1)
	}

	return x, y
}
