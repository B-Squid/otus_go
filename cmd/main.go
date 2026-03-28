package main

import (
	model "otus_go/internal"
	r "otus_go/internal/repository"
	s "otus_go/internal/service"
)

// func main() {
// 	props := model.Dialog()                                                                                                      // создадим карту с настройками игры (имена игроков, размер поля)
// 	match := model.Match{Player1: model.Player{Id: 1, Name: props["name1"]}, Player2: model.Player{Id: 1, Name: props["name2"]}} // создаем объект "партия"
// 	match.InitBoard(props)                                                                                                       // вызовем метод "партии", и создадим объект доски в соответствии с настройками
// 	match.Id = 12345                                                                                                             //TODO: инициализацию перенести в метод InitBoard (в том числе и заполнение имен игроков???)
// 	match.Chessboard.Draw(match.Player1.Name, match.Player2.Name)
// 	model.TurnDialog(match) // крутим и рисуем бесконечный цикл игрового диалога
// }

func main() {
	props := model.Dialog()

	// создадим карту с настройками игры (имена игроков, размер поля)
	match := s.Transciever(props)
	r.TurnDialog(match) // крутим и рисуем бесконечный цикл игрового диалога
}
