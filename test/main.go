package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		s := generateToken(10)
		fmt.Println(s)
		if s == "" {
			return
		}
		time.Sleep(time.Second)
	}

}

func generateToken(ln int) string {
	rand.Seed(time.Now().UnixNano())
	var token string
	for i := 0; i < ln; i++ {
		n := rand.Intn(62)
		if n < 10 {
			n += 48
		} else if n < 36 {
			n += 55
		} else {
			n += 61
		}
		token += string(rune(n))
	}
	return token
}
