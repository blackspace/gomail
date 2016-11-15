package mailstringserver

import (
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/command"
	"github.com/blackspace/gomail"
	_ "github.com/blackspace/goserver/command/help"
	"strings"
	"strconv"
	"github.com/blackspace/goserver"
	"net"
)


type MailString struct {
	Number int64
	IsReply bool
	From net.Addr
	Load string
}

func NewMailString(number int64,from net.Addr,load string) *MailString {
	return &MailString{Number:number,From:from,Load:load}
}

func init() {
	command.Commands.RegistCommand("add_mail", func(clt *client.Client, args ...string) string {
		n,_:=strconv.Atoi(args[0])
		m := NewMailString(int64(n),clt.RemoteAddr(),(strings.Join(args, " ")))
		gomail.MailBox.AddMail(m)
		return ""
	}, "")


	command.Commands.RegistCommand("add_reply", func(clt *client.Client, args ...string) string {
		n,_:=strconv.Atoi(args[0])
		m := NewMailString(int64(n),clt.RemoteAddr(),(strings.Join(args, " ")))
		m.IsReply=true
		gomail.MailBox.AddMail(m)
		return ""
	}, "")
}


var _go_mail_server = goserver.NewLineServer()

func Start() {
	_go_mail_server.Start("127.0.0.1", "5050")
}

func Stop() {
	_go_mail_server.Stop()
}



