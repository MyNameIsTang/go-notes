package demo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address1 struct {
	Type    string
	City    string
	Country string
}

type VCard1 struct {
	FirstName string
	LastName  string
	Addresses []*Address1
	Remark    string
}

func InitJson() {
	pa := &Address1{"private", "Aartselaar", "Belgium"}
	wa := &Address1{"work", "Boom", "Belgium"}
	vc := VCard1{"Jan", "Kersschot", []*Address1{pa, wa}, "none"}
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s\n", js)
	file, _ := os.OpenFile("/Users/daguang/main/go-notes/demo/vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
