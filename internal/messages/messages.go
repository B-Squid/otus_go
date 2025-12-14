package messages

type Sorter interface {
	Getid() int
}

type SdMessage struct {
	CreatedBy string
	Id        int
}

func (i SdMessage) Getid() int {
	return i.Id
}

type RichMessage struct {
	Id        int
	RouteRule string
}

func (i RichMessage) Getid() int {
	return i.Id
}

type DbMessage struct {
	Id int
}

func (i DbMessage) Getid() int {
	return i.Id
}
