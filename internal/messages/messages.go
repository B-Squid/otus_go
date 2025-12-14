package messages

type SdMessage struct {
	createdBy        string
	createdTime      string
	id               string //добавить метод  возврата ссылки на заявки по id. Типа https://<server>/api/<id>
	priority         string
	shortDescription string
	subject          string
}

type RichMessage struct {
	SdMessage
	RouteRule string
}

type DbMessage struct {
}
