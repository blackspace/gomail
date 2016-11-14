package main

import (
	"github.com/blackspace/gomail"
	"github.com/blackspace/gomail/mailclient"
	"github.com/blackspace/gomail/mailstringserver"
	"log"
	"sync"
	"time"
)

var started = make(chan int)

const TIME = 1 << 10

func _Server() {
	mailstringserver.Start()
	defer mailstringserver.Stop()
	c := 0

	started <- 1
	start := time.Now()

	for i := 0; i < TIME; i++ {
		gomail.MailBox.GetMail()
		c++
	}

	log.Println(c)
	log.Println(time.Now().Sub(start))

}

func _Client() {
	<-started
	for i := 0; i < TIME; i++ {
		mailclient.SendMailString("127.0.0.1:5050", "hello world")
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		_Server()
		wg.Done()
	}()

	go func() {
		_Client()
		wg.Done()
	}()

	wg.Wait()
}
