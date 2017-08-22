package main

import (
"fmt"
"net"
"os"
"time"
)

func sender(conn net.Conn) {
	for i := 0; i < 100; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write([]byte(words))
	}
}

func main() {
	server := "127.0.0.1:8989"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("connect success")

	go sender(conn)

	b :=make([]byte,1024)
	conn.Read(b)

	fmt.Println(string(b))

	for {
		time.Sleep(1 * 1e9)
	}
}