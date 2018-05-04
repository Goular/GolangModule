package main

import (
	"fmt"
	"encoding/hex"
)

func main() {
	src := []byte("4b2a")
	fmt.Println(src)
	encodedStr := hex.EncodeToString(src)
	fmt.Println(encodedStr)
}
