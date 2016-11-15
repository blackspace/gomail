package mailbinaryserver

import (
	"github.com/blackspace/goserver/action"
	"github.com/blackspace/goserver/client"
	"encoding/binary"
	"github.com/blackspace/gomail"
	"github.com/blackspace/goserver"
	"net"
)

type MailBinary struct {
	Number int64
	IsReply bool
	From net.Addr
	Load []byte
}

func NewMailBinary(number int64,from net.Addr,load []byte) *MailBinary {
	return &MailBinary{Number:number,From:from,Load:load}
}

const (
	ADDMAIL = 0x01
	ADDREPLY = 0x02
)

func IsAddMail(buf []byte) bool {
	if len(buf)>=18&&buf[0]==0x0A&&buf[1]==ADDMAIL {
		cl:=int64(binary.BigEndian.Uint64(buf[10:18]))
		if int64(len(buf))==18+cl {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}


func IsAddReply(buf []byte) bool {
	if len(buf)>=18&&buf[0]==0x0A&&buf[1]==ADDREPLY  {
		cl:=int64(binary.BigEndian.Uint64(buf[10:18]))
		if int64(len(buf))==18+cl {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func DoAddMail(clt *client.Client , buf []byte)  bool {
	n:=int64(binary.BigEndian.Uint64(buf[2:10]))
	m:=NewMailBinary(n,clt.RemoteAddr(),buf[18:])
	gomail.MailBox.AddMail(m)
	return true
}

func DoAddReply(clt *client.Client , buf []byte)  bool {
	n:=int64(binary.BigEndian.Uint64(buf[2:10]))
	m:=NewMailBinary(n,clt.RemoteAddr(),buf[10:len(buf)-2])
	m.IsReply=true
	gomail.MailBox.AddMail(m)
	return true
}

func init() {
	action.BinaryActions.AddAction(func(buf []byte) bool { return IsAddMail(buf)}, DoAddMail)
	action.BinaryActions.AddAction(func(buf []byte) bool { return IsAddReply(buf)}, DoAddReply)
}

var _go_mail_server = goserver.NewBinaryServer()

func Start() {
	_go_mail_server.Start("127.0.0.1", "5050")
}

func Stop() {
	_go_mail_server.Stop()
}

