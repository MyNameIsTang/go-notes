package demo

import (
	"encoding/gob"
	"log"
	"os"
)

func InitGob2() {
	pa := &Address1{"private", "Aartselaar", "Belgium"}
	wa := &Address1{"work", "Boom", "Belgium"}
	vc := VCard1{"Jan", "Kersschot", []*Address1{pa, wa}, "none"}

	file, _ := os.OpenFile("/Users/daguang/main/go-notes/demo/vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}
