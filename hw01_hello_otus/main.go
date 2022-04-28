package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	var whatYouSay = "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(whatYouSay))
}
