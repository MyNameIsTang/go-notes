package demo

import (
	"fmt"
	"html/template"
	"log"
)

func InitTemplateValidation() {

	tOk := template.New("ok")
	template.Must(tOk.Parse("/* and a comment */ some static text : {{ .Name }}"))
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("Error: ", err)
		}
	}()
	fmt.Println("The first one parsed Ok. ")
	fmt.Println("The next one ought to fail.")

	tErr := template.New("error_template")
	template.Must(tErr.Parse(" some static text {{ .Name }"))

}

// func wrap1(fn func()) {
// 	fn()
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Fatal("Error: ", err)
// 		}
// 	}()
// }
