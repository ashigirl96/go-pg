package main

import (
	"fmt"
	"github.com/ashigirl96/go-pg/pkg"
	"time"
)

func goroutines() {
	pkg.FuncName()
	say := func(s string) {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s)
		}
	}
	go say("world")
	say("hello")
}

func main() {
	goroutines()
}
