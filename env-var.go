// Environment variables are a universal mechanism for conveying
// configuration information to Unix programs.
// Let’s look at how to set, get, and list environment variables.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// To set a key/value pair, use os.Setenv.
	os.Setenv("FOO", "1")
	// To get a value for a key, use os.Getenv.
	fmt.Println("FOO:", os.Getenv("FOO"))
	// This will return an empty string if the key isn’t present in the environment.
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	// Use os.Environ to list all key/value pairs in the environment.
	// This returns a slice of strings in the form KEY=value.
	// You can strings.Split them to get the key and value.
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
