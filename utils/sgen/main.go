package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gorilla/securecookie"
)

func main() {
	fmt.Println("session auth")
	b := securecookie.GenerateRandomKey(64)
	fmt.Println(base64.StdEncoding.EncodeToString(b))

	fmt.Println("encryption key")
	b = securecookie.GenerateRandomKey(32)
	fmt.Println(base64.StdEncoding.EncodeToString(b))
}
