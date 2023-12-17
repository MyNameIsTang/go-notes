package demo

import "fmt"

type Customer2 struct {
	Name string
	Log
}

func (c *Customer2) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log.String())
}

func InitEmbedFunc2() {
	c := &Customer2{"Barak Obama", Log{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
}
