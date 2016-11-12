package mailclient

import (
	"net"
	"log"
)

func SendMail(addr string,content string) {
	conn,err:=net.Dial("tcp",addr)
	defer conn.Close()
	if err!=nil {
		log.Println(err)
	} else {
		s:="add_mail "+content+"\r\n"
		conn.Write([]byte(s))
	}

}

func SendMailOnConnect(conn net.Conn,content string) {
	s:="add_mail "+content+"\r\n"
	conn.Write([]byte(s))
}
