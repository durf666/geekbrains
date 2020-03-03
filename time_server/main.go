package main

import (
	"bufio"
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

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		handleConn(conn)

	}
}
func handleConn(c net.Conn) {
	// control := make(chan bool)
	defer c.Close()
	go readConn(c)
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
func Equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func readConn(c net.Conn) {
	for {

		bufReader := bufio.NewReader(c)
		bytes, _ := bufReader.ReadBytes('\n')
		fmt.Println(bytes)
		b := []byte{101, 120, 105, 116, 13, 10}
		if Equal(bytes, b) {
			c.Close()
			break
		}

	}
}
