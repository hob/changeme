package main

import (
	"fmt"
	"os"
)

func main() {
	println(fmt.Sprintf("Hello %s", os.Args[1]))
}
