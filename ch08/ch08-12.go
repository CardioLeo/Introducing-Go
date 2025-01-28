package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("error with listening at server()")
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("error with accepting connection at server()")
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
		fmt.Println("error at decoding message within handleServerConnection()")
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
		fmt.Println("error at calling net.Dial within client()")
		fmt.Println(err)
		return
	}

	// send the message
	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println("error at sending message from client()")
		fmt.Println(err)
	}

	defer c.Close()
}

func main() {
	var input string

	go server()
	fmt.Println("Server started")
	go client()
	fmt.Println("Client started")

	fmt.Println("starting up!")
	
	// var input string
	fmt.Scanln(&input)
}

// this consistently gets the error message:
// "dial tcp 127.0.0.1:8080: connect: connection refused"

// if the code isn't running, port :8080 says: "This site can't be reached"
// but when the code is running, Chrome tells me the following:
// "This page isn’t working
// localhost didn’t send any data.
// ERR_EMPTY_RESPONSE"

// I added print statements throughout to make it more clear exactly where the error
// was occuring, and got this back: "error at calling net.Dial within client()"
