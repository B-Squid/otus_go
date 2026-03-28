package model

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type TurnInterface interface { // ход
	Move([]string) // сходить
}

type Player struct { // игрок
	Id      int    // некий id игрока
	Name    string // имя игрока
	AutoCnt int    //счетчик автохода
}

// func (p Player) Move() string {
// 	var turn string
// 	fmt.Print("Ходит ", p.Name, "\n")
// 	fmt.Scan(&turn)
// 	fmt.Print("\033[2K\033[2A\033[2K")
// 	return turn
// }

type Turn struct { // ход игрока
	Player
	From string // откуда ходим
	To   string // куда ходим
}

func (t *Turn) GoAuto(tmpX int, tmpY int) []string {
	var str string
	var tmp []string
	var tmpLetterX, tmpLetterY rune
	var currentLetter rune = 'A'

	tmpX = rand.Intn(tmpX) + 1
	tmpY = rand.Intn(tmpY) + 1

	tmpLetterX = currentLetter + rune(tmpX) - 1
	tmpLetterY = currentLetter + rune(tmpY) - 1

	str = string(tmpLetterX) + strconv.Itoa(tmpX) + "-" + string(tmpLetterY) + strconv.Itoa(tmpY)

	fmt.Println("Автоход = ", str)

	tmp = strings.Split(strings.ToUpper(str), "-")
	return tmp
}

func (t *Turn) Dialog() (int, []string) {
	var turn, auto_turn string
	var strings_out []string

	fmt.Print("Ходит ", t.Name, "\n")
	fmt.Scan(&turn)
	//fmt.Print("\033[2K\033[2A\033[2K")
	//e2-e4
	if strings.ToLower(turn) == "retreat" {
		fmt.Println(t.Name, " сдался.")
		os.Exit(0)
	}
	if strings.ToLower(turn) == "autogame" {
		fmt.Println("Введите количество ходов:")
		fmt.Scan(&auto_turn)
	}

	strings_out = strings.Split(strings.ToUpper(turn), "-")
	tmpA_t, _ := strconv.Atoi(auto_turn)
	return tmpA_t, strings_out
}

func (t *Turn) Move(strings_out []string) {
	t.From = strings_out[0]
	t.To = strings_out[1]
}

type Chessboard struct { // доска (игровое поле)
	Id     int // некий id доски
	Size_X int // размерность доски
	Size_Y int
	Boards []string // доска
	// boardPositionX int // если досок на экране будет несколько
	// boardPositionY int
}

func (cb *Chessboard) Move(strings_out []string) {
	// смена фигур
	var iFrom, jFrom, iTo, jTo int
	var str_from, str_to string
	//вернуть позицию в cb.Boards[0], где первый элемент строки strings_out[0] = в боардс - это jFrom
	//из strings_out вырезать число (отрезать первый элемент - это и будет число) и оно = iFrom
	//аналогично с To
	str_from = strings_out[0] //откуда
	str_to = strings_out[1]   //куда
	boardHeight := len(cb.Boards)

	jFrom = strings.Index(cb.Boards[0], str_from[0:1]) //взять букву и найти на доске, вернуть позицию - это индекс столбца на доске
	iFrom, _ = strconv.Atoi(str_from[1:])              //цифра после буквы
	iFrom = boardHeight - (iFrom + 1)

	jTo = strings.Index(cb.Boards[0], str_to[0:1])
	iTo, _ = strconv.Atoi(str_to[1:])
	iTo = boardHeight - (iTo + 1)

	//если по iFrom, jFrom нет фигуры (считать символ и сравнить с клеткой черной или белой),
	// если фигуры нет, то можно вывести на экран, что фигуры нет (нужно? или пропуск хода?)
	roone := []rune(cb.Boards[iFrom])
	roone_2 := []rune(cb.Boards[iTo])

	roone_2[jTo] = roone[jFrom]
	if (iFrom+jFrom)%2 == 0 {
		roone[jFrom] = ' '
	} else {
		roone[jFrom] = '#'
	}

	cb.Boards[iFrom] = string(roone)
	cb.Boards[iTo] = string(roone_2)
}

func (cb *Chessboard) Draw(PlayerName_1 string, PlayerName_2 string) {
	for i := 0; i < cb.Size_Y+2; i++ {
		if i == 0 {
			fmt.Println(cb.Boards[i], "    ", PlayerName_1)
		} else if i == cb.Size_Y+1 {
			fmt.Println(cb.Boards[i], "    ", PlayerName_2)
		} else {
			fmt.Println(cb.Boards[i])
		}
	}
}

type Match struct { // партия (игра)
	Player1 Player
	Player2 Player
	Chessboard
	Turns []Turn
	// Win     string
	// Score   string

}

func (m *Match) InitBoard(props map[string]string) { // метод струтуры доска (первоначальная инициализация или конструктор).
	m.Size_X, _ = strconv.Atoi(props["x"]) // ширина поля
	m.Size_Y, _ = strconv.Atoi(props["y"]) // высота поля
	m.Id = 1
	/*brdArr*/ m.Boards = make([]string, m.Size_Y+2) // массив содержит буквенные и цифровые обозначения на доске

	for i := 1; i <= m.Size_X; i++ { // создаем буквы
		letterCapasitor += string(currentLetter)
		currentLetter += 1
	}

	m.Boards[0] = letterCapasitor

	for i := 1; i <= m.Size_Y+1; i++ { // главный цикл генерации игрового поля
		result = strconv.Itoa(m.Size_Y - i + 1) // номер строки

		if i == 2 { // белые пешки
			white = string(pawn)
			black = string(pawn)
		} else if i == m.Size_Y-1 { // черные пешки
			white = string(pawn + 6)
			black = string(pawn + 6)
		} else {
			white = " "
			black = "#"
		}
		var k int = 0

		for j := 1; j <= m.Size_X; j++ { // фигуры

			if i == 1 { // белые
				white = string(arrManWhite[k])
				black = string(arrManWhite[k])

				if k == len(arrManWhite)-1 {
					k = 0
				} else {
					k++
				}
			} else if i == m.Size_Y { // черные
				white = string(arrManWhite[k] + 6)
				black = string(arrManWhite[k] + 6)

				if k == len(arrManWhite)-1 {
					k = 0
				} else {
					k++
				}
			}

			if (i+j)%2 == 0 {
				result += white
			} else {
				result += black
			}
		}

		m.Boards[i] = result
		result = ""
		// fmt.Println(m.Boards[i])
	}

	m.Boards[m.Size_Y+1] = letterCapasitor

	//for i := 0; i <= y+1; i++ {
	//	fmt.Println(m.Boards[i])
	//}
}

var props = make(map[string]string) // массив настроек

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

func Dialog() map[string]string { // функция начального диалога для получения имен игроков и размера доски
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

func validate_args(args ...string) {
	for _, arg := range args {
		tst, _ := strconv.Atoi(arg)

		if tst == 0 {
			fmt.Println("Ширина или высота игрового поля не может быть равна 0. Завершение работы.")
			os.Exit(1)
		}
		if tst < 0 {
			fmt.Println("Ширина или высота игрового поля не может быть < 0. Завершение работы.")
			os.Exit(1)
		}
	}
}

// метод записи Автохода у игрока, n - счетчик для вычисления порядка хода, cnt - автоход
func (m *Match) SetPlayerAutoCnt(n int, cnt int) {
	if n%2 != 0 {
		m.Player1.AutoCnt = cnt
	} else {
		m.Player2.AutoCnt = cnt
	}
}
