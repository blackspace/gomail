package mailserver

type Mail struct {
	Number int
	From string
	Load string
}

func _NewMail(number int,from string,load string) *Mail {
	return &Mail{Number:number,From:from,Load:load}
}

type _MailBox struct {
	_channel chan *Mail
}

func _NewMailBox() *_MailBox {
	return &_MailBox{_channel: make(chan *Mail, 1<<20)}
}

func (mb *_MailBox) _AddMail(m *Mail) {
	mb._channel <- m
}

func (mb *_MailBox) GetMail() *Mail {
	return <-mb._channel
}
