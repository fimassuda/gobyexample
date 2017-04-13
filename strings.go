package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("string.split")
	var s string = "/url/path/resource1/resource2"
	stringArray := strings.Split(s, "/")
	fmt.Println(len(stringArray))
}
