package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	err := tpl.Execute(os.Stdout, fm)
	if err != nil {
		log.Fatalln(err)
	}
}
