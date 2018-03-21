package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// e1a6e33faec4522a9bb04affceaad9c3e4d7da43      sha1 for 460co8gBx4KdzLnO8Wgj + mysecret 
	
	password := []byte("mysecret")
	
	//hash := RandString(20)
	hash := []byte("460co8gBx4KdzLnO8Wgj")
	
	hp := hashPassword(hash, password)
	
	fmt.Println(hp)
	fmt.Printf("%x\n", hp)

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func hashPassword(h []byte, p []byte) []byte {
	hp := append(h, p...)
	s := sha1.New()
	s.Write(hp)
	return s.Sum(nil)
}

func RandString(n int) []byte {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return []byte(string(b))
}
