package duration_test

import (
	"reflect"
	"testing"

	"github.com/jacoelho/duration"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input   string
		want    duration.Duration
		wantErr bool
	}{
		{
			input: "P23DT23H",
			want: duration.Duration{
				Day:  23.0,
				Hour: 23.0,
			},
			wantErr: false,
		},
		{
			input: "P3Y6M4DT12H30M5S",
			want: duration.Duration{
				Year:   3.0,
				Month:  6.0,
				Day:    4.0,
				Hour:   12.0,
				Minute: 30.0,
				Second: 5.0,
			},
			wantErr: false,
		},
		{
			input: "P0.5Y",
			want: duration.Duration{
				Year: 0.5,
			},
			wantErr: false,
		},
		{
			input: "P0.5Y0.5M",
			want: duration.Duration{
				Year:  0.5,
				Month: 0.5,
			},
			wantErr: false,
		},
		{
			input: "P0,5Y0,5M",
			want: duration.Duration{
				Year:  0.5,
				Month: 0.5,
			},
			wantErr: false,
		},
		{
			input:   "P12WT12H30M5S",
			wantErr: true,
		},
		{
			input:   "P0.5S0.5M",
			wantErr: true,
		},
		{
			input:   "P5S",
			wantErr: true,
		},
		{
			input:   "P0.5A",
			wantErr: true,
		},
		{
			input:   "P0.1",
			wantErr: true,
		},
		{
			input:   "PT",
			wantErr: true,
		},
		{
			input:   "PP",
			wantErr: true,
		},
		{
			input:   "",
			wantErr: true,
		},
		{
			input:   "P0.1MM",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := duration.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
