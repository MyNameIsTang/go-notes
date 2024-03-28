package demo

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status  Status
}

type User2 struct {
	Status Status
}

func InitTwitterStatus() {
	response, _ := http.Get("http://twitter.com/users/Googland.xml")
	user := User{xml.Name{"", "user"}, Status{""}}
	reader := bufio.NewReader(response.Body)
	str, _ := reader.ReadBytes(' ')
	xml.Unmarshal(str, &user)
	fmt.Printf("status: %s", user.Status.Text)
}

func InitTwitterJSON() {
	res, _ := http.Get("http://twitter.com/users/Googland.json")
	user := User2{Status{""}}
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &user)
	fmt.Printf("status: %s", user.Status.Text)
}
