package demo

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
)

func InitDegob() {
	file, _ := os.Open("/Users/daguang/main/go-notes/demo/vcard.gob")
	defer file.Close()
	inReader := bufio.NewReader(file)
	var vc VCard1
	dnc := gob.NewDecoder(inReader)
	dnc.Decode(&vc)
	fmt.Printf("%#v", vc)
}
