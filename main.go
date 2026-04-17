package main

import (
	"fmt"

	"example.com/hello/hello" // import "hello" folder and its contents
)

func main() {
	fmt.Println(hello.Greeting) // print the variable Greeting stored in folder hello
}
