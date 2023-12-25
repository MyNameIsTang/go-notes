package demo

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() {
	filename, _ := filepath.Abs("demo/" + p.Title)
	err := ioutil.WriteFile(filename, p.Body, 0666)
	if err != nil {
		panic(err)
	}
}

func (p *Page) load(title string) (err error) {
	p.Title = title
	filename, _ := filepath.Abs("demo/" + title)
	p.Body, err = ioutil.ReadFile(filename)
	return
}

func InitWikiPart1() {
	page := &Page{
		Title: "snc.txt",
		Body:  []byte("丹尼斯啊u老师奶毒阿是asnjkda--ewqniun!"),
	}
	page.save()
	fmt.Println("save success!")
	page.load("products.txt")
	fmt.Println("load success!", page)
}
