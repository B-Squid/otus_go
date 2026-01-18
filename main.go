package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// var props []string
var props = make(map[string]string)

var black string = "#"
var white string = " "
var result string = ""
var board string = ""
var currentLetter rune = 'A'
var letterCapasitor string = " "

var king rune = '♔'
var queen rune = '♕'
var rook rune = '♖'
var bishop rune = '♗'
var knight rune = '♘'
var pawn rune = '♙'

var arrManWhite = [...]rune{'♖', '♘', '♗', '♕', '♔', '♗', '♘', '♖'}

func dialog() map[string]string {
	stdin := bufio.NewReader(os.Stdin)
	var x, y string
	var name1, name2 string

	fmt.Println("Введите размер шахматной доски в формате \"ширина[пробел]высота.\"")

	_, err := fmt.Scan(&x, &y)

	if err != nil {
		fmt.Println("Ошибка ввода размерности доски. Завершение работы.", err)
		os.Exit(1)
	}

	stdin.ReadString('\n') // очистим буфер перед новым вводом

	validate_args(x, y)

	props["x"] = x
	fmt.Println("Высота - ", props["x"])
	props["y"] = y
	fmt.Println("Ширина - ", props["y"])

	fmt.Println("Введите имя первого и второго игрока в формате \"имя1[пробел]имя2\"")

	n, err2 := fmt.Scan(&name1, &name2)

	if err2 != nil || n != 2 {
		fmt.Println("Ошибка ввода имен игроков. Завершение работы.", err2)
		os.Exit(1)
	}

	fmt.Println("Для продолжения нажмите Ввод")
	stdin.ReadString('\n') // очистим буфер на всякий случай

	props["name1"] = name1
	props["name2"] = name2

	return props
}

func create_board(props map[string]string) {
	x, _ := strconv.Atoi(props["x"])
	y, _ := strconv.Atoi(props["y"])
	name1, _ := props["name1"]
	name2, _ := props["name2"]

	for i := 1; i <= y; i++ {
		letterCapasitor = letterCapasitor + string(currentLetter)
		currentLetter += 1
	}

	for i := 1; i <= x; i++ {
		result = result + strconv.Itoa(x-i+1) // номер строки

		if i == 2 { // пешки
			white = string(pawn)
			black = string(pawn)
		} else if i == x-1 {
			white = string(pawn + 6)
			black = string(pawn + 6)
		} else {
			white = " "
			black = "#"
		}

		var k int = 0

		for j := 1; j <= y; j++ {
			if i == 1 { // фигуры
				white = string(arrManWhite[k])
				black = string(arrManWhite[k])

				if k == len(arrManWhite)-1 {
					k = 0
				} else {
					k++
				}
			} else if i == x {
				white = string(arrManWhite[k] + 6)
				black = string(arrManWhite[k] + 6)

				if k == len(arrManWhite)-1 {
					k = 0
				} else {
					k++
				}
			}

			if (i+j)%2 == 0 {
				result = result + white
			} else {
				result = result + black
			}
		}

		if i == 1 {
			result = result + "        " + name1
		} else if i == y {
			result = result + "        " + name2
		}

		result = result + "\n"
	}

	board = string(letterCapasitor) + "\n" + result + string(letterCapasitor) + "\n"
	fmt.Print(board)
}

func main() {
	props := dialog()
	create_board(props)
}

func validate_args(args ...string) {
	for _, arg := range args {
		tst, _ := strconv.Atoi(arg)

		switch {
		case tst == 0:
			fmt.Println("Ширина или высота игрового поля не может быть равна 0. Завершение работы.")
			os.Exit(1)
		case tst < 0:
			fmt.Println("Ширина или высота игрового поля не может быть < 0. Завершение работы.")
			os.Exit(1)
		}
	}
}
