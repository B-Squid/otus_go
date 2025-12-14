package service

import (
	m "canopy/internal/messages"
	"math/rand"
)

func Transciever() []m.Sorter {
	sdStr := m.SdMessage{CreatedBy: "a", Id: randInt()}
	rmStr := m.RichMessage{RouteRule: "x", Id: randInt()}
	dmStr := m.DbMessage{Id: randInt()}
	myarr := []m.Sorter{sdStr, rmStr, dmStr}

	return myarr
}

func randInt() int {
	return 1 + rand.Intn(1000-1)
}
