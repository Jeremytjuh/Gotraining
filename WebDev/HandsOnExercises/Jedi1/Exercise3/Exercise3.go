package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

type restaurants []restaurant

type restaurant struct {
	Name string
	Menu menu
}

type menu struct {
	Breakfast meal
	Lunch     meal
	Dinner    meal
}

type meal struct {
	Dish        string
	Price       float64
	ServedFrom  time.Time
	ServedUntil time.Time
}

const timeParser = "03:04"

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	BreakFrom, _ := time.Parse(time.Kitchen, "06:00")
	BreakUntil, _ := time.Parse(time.Kitchen, "10:00")

	LunchFrom, _ := time.Parse(time.Kitchen, "12:00")
	LunchUntil, _ := time.Parse(timeParser, "14:00")

	DinnerFrom, _ := time.Parse(timeParser, "18:00")
	DinnerUntil, _ := time.Parse(timeParser, "20:30")

	pass := restaurants{
		restaurant{
			Name: "Chegg E Cheese",
			Menu: menu{
				Breakfast: meal{
					Dish:        "Egg",
					Price:       4.20,
					ServedFrom:  BreakFrom,
					ServedUntil: BreakUntil,
				},
				Lunch: meal{
					Dish:        "Eggxtra",
					Price:       6.90,
					ServedFrom:  LunchFrom,
					ServedUntil: LunchUntil,
				},
				Dinner: meal{
					Dish:        "Eggxtreme",
					Price:       10.00,
					ServedFrom:  DinnerFrom,
					ServedUntil: DinnerUntil,
				},
			},
		},

		restaurant{
			Name: "Wen dies",
			Menu: menu{
				Breakfast: meal{
					Dish:        "Hamburger",
					Price:       4.20,
					ServedFrom:  BreakFrom,
					ServedUntil: BreakUntil,
				},
				Lunch: meal{
					Dish:        "Hamburger",
					Price:       6.90,
					ServedFrom:  LunchFrom,
					ServedUntil: LunchUntil,
				},
				Dinner: meal{
					Dish:        "Hamburger",
					Price:       10.00,
					ServedFrom:  DinnerFrom,
					ServedUntil: DinnerUntil,
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", pass)
	if err != nil {
		log.Fatalln("An error has occured")
	}
}
