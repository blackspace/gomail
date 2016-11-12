package mailserver

import (
	"github.com/blackspace/gotask/runnable_pool/channel"
	"github.com/blackspace/goserver"
	"github.com/blackspace/goserver/command"
	_ "github.com/blackspace/goserver/command/help"
	"github.com/blackspace/goserver/client"
)

var runnable_pool=channel.NewRunnablePoolWithChannel()

func init() {
	runnable_pool.Run()

	command.Commands.RegistCommand("add_mail",func(clt *client.Client,args ...string) string {
		m:=NewMail()
		m.From=clt.RemoteAddr()
		m.Load =args
		MailBox.AddMail(m)
		return ""
	},"Add a mail")
}


var _gomail_server =goserver.NewServer()
func Start() {
	_gomail_server.Start("127.0.0.1","5050")
}

func Stop() {
	_gomail_server.Stop()
}

var MailBox=NewMailBox()

