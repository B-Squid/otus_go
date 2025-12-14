package repository

import (
	m "canopy/internal/messages"
	"fmt"
)

var SliceForSd []m.SdMessage
var SliceForRich []m.RichMessage
var SliceForDb []m.DbMessage

func Reciever(s []m.Sorter) {

	for i := range s {
		switch v := s[i].(type) {
		case m.SdMessage:
			fmt.Println("Processing SdMessage, Id =", v.Id)
			SliceForSd = append(SliceForSd, s[i].(m.SdMessage))
		case m.RichMessage:
			fmt.Println("Processing RichMessage, Id =", v.Id)
			SliceForRich = append(SliceForRich, s[i].(m.RichMessage))
		case m.DbMessage:
			fmt.Println("Processing DbMessage, Id =", v.Id)
			SliceForDb = append(SliceForDb, s[i].(m.DbMessage))
		default:
			fmt.Println("Unknown message type :(")
		}
	}
}
