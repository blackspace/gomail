package main

import (
	"github.com/blackspace/gomail/mailserver"
	"log"

	"time"
)


func main() {
	mailserver.Start()
	defer mailserver.Stop()
	c:=0
	start:=time.Now()

	for i:=0;i<100000;i++{
		mailserver.MailBox.GetMail()
		c++
	}


	log.Println(c)
	log.Println(time.Now().Sub(start))

}