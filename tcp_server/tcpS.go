package main

import (
	"net"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	ln, _ := net.Listen("tcp", ":8082")
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message from client :", message)
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage))
	}
}
