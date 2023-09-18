package main

import (
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type Result struct {
	Result int `json:"result"`
}

type App struct {
	Arrays [][]string
}

func (app *App) Handler(w http.ResponseWriter, r *http.Request) {
	i := 0

	for b := 0; b < 50000; b++ {
		array := app.Arrays[b%len(app.Arrays)]

		for j := range array {
			if array[j] == "qw2" {
				i++

				break
			}
		}
	}

	time.Sleep(10 * time.Millisecond)

	for b := 0; b < 50000; b++ {
		array := app.Arrays[b%len(app.Arrays)]

		for j := range array {
			if array[j] == "qw5" {
				i++

				break
			}
		}
	}

	time.Sleep(10 * time.Millisecond)

	for b := 0; b < 50000; b++ {
		array := app.Arrays[b%len(app.Arrays)]

		for j := range array {
			if array[j] == "qw8" {
				i++

				break
			}
		}
	}

	err := json.NewEncoder(w).Encode(Result{
		Result: i,
	})

	if err != nil {
		panic(err)
	}
}

func main() {
	app := App{
		Arrays: make([][]string, 0),
	}

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			app.Arrays = append(app.Arrays, []string{"qw1", "qw2", "qw3", "qw4", "qw5"})
		} else {
			app.Arrays = append(app.Arrays, []string{"qw5", "qw4", "qw3", "qw2", "qw1"})
		}
	}

	http.HandleFunc("/", app.Handler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

	println(1)
}
