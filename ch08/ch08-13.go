package main

import (
	"fmt"
	"net/http"
	"io"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("calling res.Header().Set()")
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Println("calling io.WriteString()")
	io.WriteString(
		res,
		`<!DOCTYPE html>
		<html>
			<head>
				<title>Hello, World</title>
			</head>
			<body>
				<p>Hello, World!</p>
			</body>
		</html>`,
	)
}

/*
http.Handle(
	"/assets/",
	http.StripPrefix(
		"/assets/",
		http.FileServer(http.Dir("assets"),)
	),
)
// code from bottom of page 77
*/

func main() {
	fmt.Println("calling http.HandleFunc()")
	http.HandleFunc("/hello", hello)
	fmt.Println("calling http.ListenAndServe()")
	http.ListenAndServe("9000", nil)
	fmt.Println("exiting")
}

// currently the only output is this:

// calling http.HandleFunc()
// calling http.ListenAndServe()
// exiting

// which means it isn't calling hello - which would have made sense if I had looked closer at the code
