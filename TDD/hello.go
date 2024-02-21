package main

import (
	"fmt"
)

const (
	enHelloPrefix = "Hello, "
	ptHelloPrefix = "Ola, "
	jpHelloPrefix = "こんにちは, "
	portugues     = "Portugues"
	japanese      = "Japanese"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portugues:
		prefix = ptHelloPrefix
	case japanese:
		prefix = jpHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
