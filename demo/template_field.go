package demo

import (
	"fmt"
	"html/template"
	"os"
)

type Person4 struct {
	Name                string
	nonExportedAgeField string
}

func InitTemplateField() {
	t := template.New("hello")
	t, _ = t.Parse("hello {{.Name}}!")
	p := &Person4{"Tom", "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
