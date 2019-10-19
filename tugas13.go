package main

import "fmt"
import "encoding/base64"

import "crypto/sha1"

func main() {
	//encoding base64
	var text string
	fmt.Print("masukan nama : ")
	fmt.Scanln(&text)

	var encodeString = base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println("encoding nama :", encodeString)

	//enkripsi sha1
	var text2 = encodeString
	var sha = sha1.New()
	sha.Write([]byte(text2))
	var enkripsi = sha.Sum(nil)
	var stringenkripsi = fmt.Sprintf("%x", enkripsi)

	fmt.Println("enkripsi nama : ", stringenkripsi)

}
