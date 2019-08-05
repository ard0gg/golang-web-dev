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

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m1 := menu{
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

	m2 := menu{
		meal{
			Name: "Breakfast",
			Items: []menuItem{
				menuItem{
					Name:        "Full Scramble 2",
					Description: "Eggs, Bacon, Toast, Hasbrowns",
					Price:       13.50,
				},
				menuItem{
					Name:        "Half Scramble 2",
					Description: "Eggs, Bacon",
					Price:       10.40,
				},
			},
		},
		meal{
			Name: "Lunch",
			Items: []menuItem{
				menuItem{
					Name:        "sandwhich 2",
					Description: "PB&J",
					Price:       15.32,
				},
			},
		},
	}

	r := restaurants{
		restaurant{
			Name: "Arby's",
			Menu: m1,
		},
		restaurant{
			Name: "AppleBee's",
			Menu: m2,
		},
	}
	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
