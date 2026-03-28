package repository

import (
	model "otus_go/internal"
)

func TurnDialog(match model.Match) { // передадим в функцию объект Match, содержащий поля с именами игроков
	var n int = 1
	var coordinates []string
	var autocount int

	for {
		if n%2 != 0 {
			match.Turns = append(match.Turns, model.Turn{Player: model.Player{Id: match.Player1.Id, Name: match.Player1.Name, AutoCnt: match.Player1.AutoCnt}})
		} else {
			match.Turns = append(match.Turns, model.Turn{Player: model.Player{Id: match.Player2.Id, Name: match.Player2.Name, AutoCnt: match.Player2.AutoCnt}})
		}

		if match.Turns[len(match.Turns)-1].AutoCnt == 0 { //если в предыдущих ходов не было Автохода, то запросим ход
			autocount, coordinates = match.Turns[len(match.Turns)-1].Dialog( /*match*/ )
			match.SetPlayerAutoCnt(n, autocount)                //запомним у того или иного игрока
			match.Turns[len(match.Turns)-1].AutoCnt = autocount //запомним у Хода
		}

		if match.Turns[len(match.Turns)-1].AutoCnt > 0 { // проверить, был введен Автоход или с прошлых ходов остался, то вызовем Автоход
			// вызвать автоход
			coordinates = match.Turns[len(match.Turns)-1].GoAuto(match.Size_X, match.Size_Y)
			match.SetPlayerAutoCnt(n, match.Turns[len(match.Turns)-1].AutoCnt-1) //уменьшим значение у того или иного игрока
		}

		match.Turns[len(match.Turns)-1].Move(coordinates)
		match.Chessboard.Move(coordinates)
		match.Chessboard.Draw(match.Player1.Name, match.Player2.Name)
		n++
	}
}
