package gomail

type Mail interface {}

type _MailBox struct {
	_channel chan Mail
}

func NewMailBox() *_MailBox {
	return &_MailBox{_channel: make(chan Mail, 1<<20)}
}

func (mb *_MailBox) AddMail(m Mail) {
	mb._channel <- m
}


func (mb *_MailBox) GetMail() Mail {
	return <-mb._channel
}

var MailBox = NewMailBox()




