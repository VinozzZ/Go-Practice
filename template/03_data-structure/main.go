package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type StructValue struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	/*
		Slice
	*/
	// sages := []string{"Gandi", "MLK", "Buddha"}
	// err := tpl.Execute(os.Stdout, sages)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	/*
		Map
	*/
	// mapValue := map[string]string{"Gandi": "human", "MLK": "god", "Buddha": "what"}
	// err := tpl.Execute(os.Stdout, mapValue)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	/*
		Struct
	*/
	structValue := StructValue{
		Name:  "Sofia",
		Motto: "time",
	}
	err := tpl.Execute(os.Stdout, structValue)
	if err != nil {
		log.Fatalln(err)
	}

}
