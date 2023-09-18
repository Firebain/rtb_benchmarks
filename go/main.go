package main

import "time"

func main() {
	// runOld()
	runNew()
}

func testable(app App) Result {
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

	return Result{
		Result: i,
	}
}
