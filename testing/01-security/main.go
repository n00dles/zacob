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
	hp := append(h, p...)  									// concatenate the 2 byte arrays
	s := sha1.New() 										// create a new hash
	s.Write(hp)												// hash hp
	return s.Sum(nil)										// return it
}

func RandString(n int) []byte {
	b := make([]rune, n)									// create a slice b of type rune of length n
	for i := range b {										// loop for lenght n
		b[i] = letterRunes[rand.Intn(len(letterRunes))]		// get a randon letter from letterRunes and add to b[]
	}
	return []byte(string(b))								// return it
}
