package main

import "testing"

func BenchmarkTest(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		testable(app)
	}
}
