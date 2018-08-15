package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// ParseFiles can parse a specified file
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("two.ghtml", "three.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// ExecuteTemplate is used for multiple files in one pointer to specify which template to execute
	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// ParseGlob can parse multiple files
	tpl, err = template.ParseGlob("template/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

}
