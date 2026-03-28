package service

import (
	model "otus_go/internal"
)

func Transciever(props map[string]string) model.Match {
	match := model.Match{Player1: model.Player{Id: 1, Name: props["name1"]}, Player2: model.Player{Id: 1, Name: props["name2"]}} // создаем объект "партия"
	match.InitBoard(props)                                                                                                       // вызовем метод "партии", и создадим объект доски в соответствии с настройками
	match.Id = 12345
	match.Chessboard.Draw(match.Player1.Name, match.Player2.Name)

	return match
}
