package messages

import "sync"

type Sorter interface {
	Getid() int
}

type SdMessage struct {
	SdmLocker sync.RWMutex
	CreatedBy string
	Id        int
}

func (i SdMessage) Getid() int {
	return i.Id
}

type RichMessage struct {
	RmLocker  sync.RWMutex
	CreatedBy string
	Id        int
	RouteRule string
}

func (i RichMessage) Getid() int {
	return i.Id
}

type DbMessage struct {
	DmLocker  sync.RWMutex
	CreatedBy string
	Id        int
}

func (i DbMessage) Getid() int {
	return i.Id
}
