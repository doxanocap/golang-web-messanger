package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	f, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	_, err2 := f.WriteString("qweqeqwrq")
	if err2 != nil {
		panic(err2)
	}
}
