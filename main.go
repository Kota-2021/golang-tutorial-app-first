package main

import (
	"html/template"
	"net/http"
)

type person struct {
	First  string
	Last   string
	Saying string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", indexShow)
	http.ListenAndServe(":8080", nil)
}

func indexShow(w http.ResponseWriter, r *http.Request) {

	//io.WriteString(w, "some text hello world")

	p1 := person{
		First:  "HTTPserver",
		Last:   "Start",
		Saying: "powered by Golang",
	}

	p2 := person{
		First:  "Hello",
		Last:   "World !",
		Saying: "Have a nice day. (≧▽≦)ノ",
	}

	xp := []person{p1, p2}

	tpl.ExecuteTemplate(w, "index.gohtml", xp)
}
