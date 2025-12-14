package service

import (
	m "canopy/internal/messages"
	"fmt"
	"math/rand"
	"time"
)

func Transciever(ch chan []m.Sorter, i int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Start transciever ", i)
	sdStr := m.SdMessage{CreatedBy: "a", Id: rand1000Int()}
	rmStr := m.RichMessage{RouteRule: "x", Id: rand1000Int()}
	dmStr := m.DbMessage{Id: rand1000Int()}
	myarr := []m.Sorter{sdStr, rmStr, dmStr}

	ch <- myarr
	//close(ch)
}

func rand1000Int() int {
	return 1 + rand.Intn(1000-1)
}

func Rand10Int() int {
	return 1 + rand.Intn(10-1)
}
