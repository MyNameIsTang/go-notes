package demo

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func InitHashSha1() {
	hasher := sha1.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("result: %x\n", hasher.Sum(b))
	fmt.Printf("result: %d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("result: %x\n", checksum)
}
