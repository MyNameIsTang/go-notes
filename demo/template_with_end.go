package demo

import (
	"html/template"
	"os"
)

func InitTemplateWithEnd() {
	t := template.New("test")
	// t, _ = t.Parse("{{with `hello`}}-oo-{{.}}-hh-{{end}}!\n")
	// t.Execute(os.Stdout, nil)

	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n")
	t.Execute(os.Stdout, nil)
}
