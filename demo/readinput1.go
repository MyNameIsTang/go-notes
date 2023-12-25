package demo

import "fmt"

var (
	firstName, lastName, s1 string
	i1                      int
	f1                      float32
	input                   = "56.12 / 5212 / Go"
	format                  = "%f / %d / %s"
)

func InitReadinput1() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f1, &i1, &s1)
	fmt.Println("From the string we read: ", f1, i1, s1)
}
