package duration_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/jacoelho/duration"
)

func TestDuration_UnmarshalText(t *testing.T) {
	var subject struct {
		Duration duration.Duration `json:"duration,omitempty"`
	}

	err := json.Unmarshal([]byte(`{"duration":"P3Y6M4DT12H30M5S"}`), &subject)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	want := duration.Duration{
		Year:   3.0,
		Month:  6.0,
		Day:    4.0,
		Hour:   12.0,
		Minute: 30.0,
		Second: 5.0,
	}

	if !reflect.DeepEqual(want, subject.Duration) {
		t.Fatalf("expected %s, got %s", want, subject.Duration)
	}
}

func TestDuration_MarshalText(t *testing.T) {
	subject := struct {
		Duration duration.Duration `json:"duration,omitempty"`
	}{
		Duration: duration.Duration{
			Year:   3.0,
			Month:  6.0,
			Day:    4.0,
			Hour:   12.0,
			Minute: 30.0,
			Second: 5.0,
		},
	}

	data, err := json.Marshal(subject)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	want := `{"duration":"P3Y6M4DT12H30M5S"}`
	if got := string(data); got != want {
		t.Fatalf("expected %s, got %s", want, got)
	}
}
