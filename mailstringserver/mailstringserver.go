package mailstringserver

import (
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/command"
	"github.com/blackspace/gomail"
	_ "github.com/blackspace/goserver/command/help"
	"strings"
	"strconv"
	"github.com/blackspace/goserver"
)


type MailString struct {
	Number int64
	IsReply bool
	From string
	Load string
}

func NewMailString(number int64,from string,load string) *MailString {
	return &MailString{Number:number,From:from,Load:load}
}

func init() {
	command.Commands.RegistCommand("add_mail", func(clt *client.Client, args ...string) string {
		n,_:=strconv.Atoi(args[0])
		m := NewMailString(int64(n),clt.RemoteAddr().String(),(strings.Join(args, " ")))
		gomail.MailBox.AddMail(m)
		return ""
	}, "")


	command.Commands.RegistCommand("add_reply", func(clt *client.Client, args ...string) string {
		n,_:=strconv.Atoi(args[0])
		m := NewMailString(int64(n),clt.RemoteAddr().String(),(strings.Join(args, " ")))
		m.IsReply=true
		gomail.MailBox.AddMail(m)
		return ""
	}, "")
}


var _go_mail_server = goserver.NewServer(goserver.LINEMODE)

func Start() {
	_go_mail_server.Start("127.0.0.1", "5050")
}

func Stop() {
	_go_mail_server.Stop()
}



