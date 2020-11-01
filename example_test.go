package duration_test

import (
	"fmt"

	"github.com/jacoelho/duration"
)

func ExampleParse() {
	data, _ := duration.Parse("P3Y6M4DT12H30M5S")
	fmt.Printf("%#v\n", data)
	// Output: duration.Duration{Year:3, Month:6, Day:4, Hour:12, Minute:30, Second:5}
}
