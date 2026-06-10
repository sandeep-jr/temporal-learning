package main

import (
	greeting "101/samples/greeting"
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	greeting := greeting.GreetSomeone(name)
	fmt.Println(greeting)
}
