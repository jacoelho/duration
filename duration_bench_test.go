package duration_test

import (
	"testing"

	"github.com/jacoelho/duration"
)

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, _ = duration.Parse("P3Y6M4DT12H30M5S")
	}
}
