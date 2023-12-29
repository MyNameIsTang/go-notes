package demo

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func InitXml() {
	var t xml.Token
	var err error

	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"

	inputReader1 := strings.NewReader(input)
	p := xml.NewDecoder(inputReader1)
	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
			fmt.Println("啥都没有")
		}
	}
}
