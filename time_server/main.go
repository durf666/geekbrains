package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	control := make(chan bool)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		handleConn(conn, control)
		z := <-control
		fmt.Println(z)
		if z {
			time.Sleep(1 * time.Second)
			break
		}
	}
}
func handleConn(c net.Conn, control chan bool) {

	defer c.Close()
	go readConn(c, control)
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func readConn(c net.Conn, control chan bool) {
	for {
		bufReader := bufio.NewReader(c)
		inputBytes, _ := bufReader.ReadBytes('\n')
		fmt.Println(inputBytes)
		b := []byte{101, 120, 105, 116, 13, 10}
		fmt.Println("111")
		if bytes.Equal(inputBytes, b) {
			c.Close()
			control <- true
			break
		}

	}
}
