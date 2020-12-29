package main

import "fmt"

const prefix = "hello, "

// Hello echo hello, world
func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
