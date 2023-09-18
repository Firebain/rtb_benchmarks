package main

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func (app *App) Handler(ctx *fiber.Ctx) error {
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

	return ctx.JSON(Result{
		Result: i,
	})
}

func runNew() {
	api := fiber.New(fiber.Config{
		Prefork:     true,
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

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

	api.Get("/", app.Handler)

	if err := api.Listen(":8080"); err != nil {
		panic(err)
	}
}
