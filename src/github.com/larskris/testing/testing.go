package main

import (
	"fmt"
)

func say(name string) (answer, goodbye string) {
	answer = "Hello " + name
	goodbye = "Now leave, " + name
	return
}

func main() {
	a, g := say("Kristian")
	fmt.Println(a)
	fmt.Println(g)
}
