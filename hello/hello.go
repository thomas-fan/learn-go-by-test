package main

import "fmt"

const englishPrefix = "hello, "
const spanishPrefix = "hola, "
const frenchPrefix = "bonjour, "

// Hello echo hello, world
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchPrefix
	case "spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
