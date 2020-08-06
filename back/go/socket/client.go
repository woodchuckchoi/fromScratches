package main

import (
	"net"
	"fmt"
	_ "io"
	"os"
	"bufio"
)

func main() {
	conn, err := net.Dial("tcp", ":7777")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for scanner.Scan() {
			conn.Write([]byte(scanner.Text()))
		}
	}()

	go func() {
		recv := make([]byte, 4096)
		for {
			n, err := conn.Read(recv)
			if err != nil {
				panic(err)
				os.Exit(1)
			}
			fmt.Println(string(recv[:n]))
		}
	}()

	for {
	}
}
