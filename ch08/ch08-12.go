package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	defer c.Close()
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:8080") // book uses port ":9999";
	if err != nil {
		fmt.Println(err)
		return
	}

	// send the message
	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	defer c.Close()
}

func main() {
	go server()
	fmt.Println("Server started")
	go client()
	fmt.Println("Client started")

	fmt.Println("starting up!")
	var input string
	fmt.Scanln(&input)
}

// this consistently gets the error message:
// "dial tcp 127.0.0.1:8080: connect: connection refused"
