package main

import "net"

import "bufio"

import "os"

import "fmt"

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8082")

	for {
		recivetest, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		// send socket
		fmt.Fprintf(conn, recivetest)
		// receive from socket
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("New message : ", message)
	}
}
