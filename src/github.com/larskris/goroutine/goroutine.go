package main 

import (
	"fmt"
	"time"
)


func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
		fmt.Println(time.Now().Second())
	}
}

func main() {
	go say("also running")
	say("running") 
}