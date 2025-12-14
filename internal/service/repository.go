package service

import (
	m "canopy/internal/messages"
	"fmt"
	"time"
)

var SliceForSd []m.SdMessage
var SliceForRich []m.RichMessage
var SliceForDb []m.DbMessage

func Reciever(ch <-chan []m.Sorter, i int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Start reciever ", i)
	for {
		s := <-ch
		for i := range s {
			switch v := s[i].(type) {
			case m.SdMessage:
				fmt.Println("Processing SdMessage, Id =", v.Id)
				v.SdmLocker.Lock()
				SliceForSd = append(SliceForSd, s[i].(m.SdMessage))
				v.SdmLocker.Unlock()
			case m.RichMessage:
				fmt.Println("Processing RichMessage, Id =", v.Id)
				v.RmLocker.Lock()
				SliceForRich = append(SliceForRich, s[i].(m.RichMessage))
				v.RmLocker.Unlock()
			case m.DbMessage:
				fmt.Println("Processing DbMessage, Id =", v.Id)
				v.DmLocker.Lock()
				SliceForDb = append(SliceForDb, s[i].(m.DbMessage))
				v.DmLocker.Unlock()
			default:
				fmt.Println("Unknown message type :(")
			}
		}
	}
}
