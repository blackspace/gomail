package mailclient

import (
	"log"
	"net"
	"strings"
	"strconv"
)

var _number=0

func SendMail(addr string, content string) {
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		log.Println(err)
	} else {
		s := strings.Join([]string{"add_mail",strconv.Itoa(_number),content,"\r\n"}," ")
		conn.Write([]byte(s))
		_number++
	}

}

func SendMailOnConnect(conn net.Conn, content string) {
	s := strings.Join([]string{"add_mail",strconv.Itoa(_number),content,"\r\n"}," ")
	conn.Write([]byte(s))
	_number++
}
