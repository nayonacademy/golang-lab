package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func getCode(s string) string{
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main(){
	c := getCode("test@gmail.com")
	fmt.Println(c)
	c = getCode("test@gmail.com")
	fmt.Println(c)
}
