package main

import (
	"fmt"
	"html/template"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {

	templateNames := [5]string{"welcome", "thanks", "sorry", "list", "form"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("templates/layout.html", "templates/"+name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func main() {
	loadTemplates()
}
