package main

import (
	"fmt"
)
import "os"

func main() {
	fmt.Println("a")
	args := os.Args
	fmt.Println(args[0])
}
