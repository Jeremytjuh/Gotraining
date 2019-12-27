package randomrune

import (
	"math/rand"
	"time"
)

// RandomIndexSlice Wowie
func RandomIndexSlice(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyz"
	password := make([]byte, length)
	for range password {
		password = append(password, chars[rand.Intn(len(chars)-1)])
	}
	return string(password)
}

// RandomASCII IDTrust dokkements
func RandomASCII(length int) string {
	rand.Seed(time.Now().UnixNano())
	var password []rune
	for i := 0; i < length; i++ {
		getal := rand.Intn((125))
		for getal < 33 {
			getal = rand.Intn(125)
		}
		password = append(password, rune(getal))
	}
	return string(password)
}
