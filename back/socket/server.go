package main

import (
	_ "io"
	"fmt"
	"net"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func HandleConn(conn net.Conn) {
	buff := make([]byte, 4096)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Connection Lost!")
			return
		}
		fmt.Println(string(buff[:n]))
		_, err = conn.Write(buff[:n])
		if err != nil {
			fmt.Println("Connection Lost!")
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", ":7777")
	HandleErr(err)
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()
		go HandleConn(conn)
	}
}

