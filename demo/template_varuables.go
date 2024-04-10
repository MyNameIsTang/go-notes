package demo

import (
	"html/template"
	"os"
)

func InitTempalteVariables() {
	t := template.New("test")
	t = template.Must(t.Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)

	t = template.New("test1")
	t = template.Must(t.Parse("{{with $x3 := `hole`}}{{$x3}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)

	t = template.New("test2")
	t = template.Must(t.Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
}
