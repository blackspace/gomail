package mailserver

import (
	"github.com/blackspace/goserver"
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/command"
	_ "github.com/blackspace/goserver/command/help"
	"github.com/blackspace/gotask/runnable_pool/channel"
	"strings"
	"strconv"
)

var runnable_pool = channel.NewRunnablePoolWithChannel()

func init() {
	runnable_pool.Run()

	command.Commands.RegistCommand("add_mail", func(clt *client.Client, args ...string) string {
		n,_:=strconv.Atoi(args[0])
		m := _NewMail(n,clt.RemoteAddr().String(),(strings.Join(args, " ")))
		MailBox._AddMail(m)
		return ""
	}, "")


	command.Commands.RegistCommand("add_reply", func(clt *client.Client, args ...string) string {
		n,_:=strconv.Atoi(args[0])
		m := _NewMail(n,clt.RemoteAddr().String(),(strings.Join(args, " ")))
		m.IsReply=true
		MailBox._AddMail(m)
		return ""
	}, "")
}

var _gomail_server = goserver.NewServer()

func Start() {
	_gomail_server.Start("127.0.0.1", "5050")
}

func Stop() {
	_gomail_server.Stop()
}

var MailBox = _NewMailBox()


