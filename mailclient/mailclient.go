package mailclient

import (
	"log"
	"net"
	"strings"
	"strconv"
	"github.com/blackspace/gomail/mailbinaryserver"
	"encoding/binary"
)

var _number=0

func SendMailString(addr string, content string) {
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

func SendReplyString(addr string, number int,content string) {
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		log.Println(err)
	} else {
		s := strings.Join([]string{"add_reply",strconv.Itoa(number),content,"\r\n"}," ")
		conn.Write([]byte(s))
	}

}

func SendReplyStringOnConnect(conn net.Conn, number int,content string) {
	s := strings.Join([]string{"add_repy",strconv.Itoa(number),content,"\r\n"}," ")
	conn.Write([]byte(s))
}


func SendMailStringOnConnect(conn net.Conn, content string) {
	s := strings.Join([]string{"add_mail",strconv.Itoa(_number),content,"\r\n"}," ")
	conn.Write([]byte(s))
	_number++
}

func SendMailBinary(addr string, content []byte) {
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		log.Println(err)
	} else {
		nb :=make([]byte,8)
		binary.BigEndian.PutUint64(nb,uint64(_number))
		clb:=make([]byte,8)
		binary.BigEndian.PutUint64(clb,uint64(len(content)))
		s := append([]byte(nil),0x0a)
		s= append(s,mailbinaryserver.ADDMAIL)
		s=append(s, nb...)
		s=append(s,clb...)
		s=append(s,content...)
		conn.Write([]byte(s))
		_number++
	}

}

func SendMailBinaryOnConnect(conn net.Conn, content []byte) {
	nb :=make([]byte,8)
	clb:=make([]byte,8)
	binary.BigEndian.PutUint64(nb,uint64(_number))
	binary.BigEndian.PutUint64(clb,uint64(len(content)))
	s := append([]byte(nil),0x0a)
	s= append(s,mailbinaryserver.ADDMAIL)
	s=append(s, nb...)
	s=append(s,clb...)
	s=append(s,content...)
	conn.Write([]byte(s))
	_number++
}

func SendReplyBinary(addr string, number int,content []byte) {
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		log.Println(err)
	} else {
		nb :=make([]byte,8)
		binary.BigEndian.PutUint64(nb,uint64(_number))
		clb:=make([]byte,8)
		binary.BigEndian.PutUint64(clb,uint64(len(content)))
		s := append([]byte(nil),0x0a)
		s= append(s,mailbinaryserver.ADDMAIL)
		s=append(s, nb...)
		s=append(s,clb...)
		s=append(s,content...)
		conn.Write([]byte(s))
	}

}

func SendReplyBinaryOnConnect(conn net.Conn, number int,content []byte) {
	nb :=make([]byte,8)
	binary.BigEndian.PutUint64(nb,uint64(_number))
	clb:=make([]byte,8)
	binary.BigEndian.PutUint64(clb,uint64(len(content)))
	s := append([]byte(nil),0x0a)
	s= append(s,mailbinaryserver.ADDMAIL)
	s=append(s, nb...)
	s=append(s,clb...)
	s=append(s,content...)
	conn.Write([]byte(s))
}