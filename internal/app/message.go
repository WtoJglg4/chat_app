package chat

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

func (mes *Message) String() string {
	return mes.Author + " says " + mes.Body

}
