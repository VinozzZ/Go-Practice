package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type CAHotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

func isInRegions(s string) bool {
	return s == "Southern" || s == "Central" || s == "Northern"
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	hotels := []CAHotel{
		CAHotel{"wala", "1234", "LA", 1234, "Southern"},
		CAHotel{"pata", "1234", "SA", 4324, "Central"},
		CAHotel{"tttt", "1234", "PA", 32432, "Northern"},
		CAHotel{"zzz", "1234", "TA", 8768, "Southern"},
	}
	for idx, hotel := range hotels {
		if strings.Contains(hotel.Region, "Southern") {
			fmt.Println(hotels[:idx])
			hotels = append(hotels[:idx])
		}

	}
	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
