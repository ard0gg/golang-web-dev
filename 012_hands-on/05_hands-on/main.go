package main

import (
	"html/template"
	"log"
	"os"
)

type menu []meal

type meal struct {
	Name  string
	Items []menuItem
}

type menuItem struct {
	Name, Description string
	Price             float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		meal{
			Name: "Breakfast",
			Items: []menuItem{
				menuItem{
					Name:        "Full Scramble",
					Description: "Eggs, Bacon, Toast, Hasbrowns",
					Price:       13.50,
				},
				menuItem{
					Name:        "Half Scramble",
					Description: "Eggs, Bacon",
					Price:       10.40,
				},
			},
		},
		meal{
			Name: "Lunch",
			Items: []menuItem{
				menuItem{
					Name:        "sandwhich",
					Description: "PB&J",
					Price:       15.32,
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
