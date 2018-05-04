package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	arr := []byte{51, 46, 52}
	bits := binary.BigEndian.Uint32(arr)
	value := math.Float32frombits(bits)
	fmt.Println(value)
}
