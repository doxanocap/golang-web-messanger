package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(string(time.Now().Format("02.01.2006, 15:04:05")))
}
