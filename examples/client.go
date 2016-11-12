package main

import (
	"github.com/blackspace/gomail/mailclient"
	"net"
	"log"
)

func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:5050")
	defer func(){
		conn.Close()
	}()
	if err==nil {
		for i:=0;i<100000;i++ {
			mailclient.SendMailOnConnect(conn,"hello")
		}
	} else {
		log.Println(err)
	}
}
