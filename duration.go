package duration

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Duration represents a iso8601 duration
type Duration struct {
	Year   float64
	Month  float64
	Day    float64
	Hour   float64
	Minute float64
	Second float64
}

// ErrParsing represents a parsing error
var ErrParsing = errors.New("parsing error")

// Parse parses a duration in iso8601 format
func Parse(input string) (Duration, error) {
	if strings.Contains(input, ",") {
		input = strings.ReplaceAll(input, ",", ".")
	}

	scanner := newScanner(strings.NewReader(input))
	scanner.Split(scanUnits)

	var (
		timePortion  bool
		duration     Duration
		field        = -1
		currentField int
	)
	for scanner.Scan() {
		token := scanner.Text()
		currentField = fieldIndex(token, timePortion)

		switch token {
		case "P":
			//
		case "T":
			timePortion = true
		default:
			if len(token) < 2 {
				return Duration{}, fmt.Errorf("unexpected field `%s`: %w", token, ErrParsing)
			}

			currentField = fieldIndex(token[len(token)-1:], timePortion)

			value, err := strconv.ParseFloat(token[:len(token)-1], 64)
			if err != nil {
				return Duration{}, fmt.Errorf("invalid decimal `%s`: %w", token, ErrParsing)
			}

			switch currentField {
			case 1:
				duration.Year = value
			case 2:
				duration.Month = value
			case 3:
				duration.Day = value
			case 5:
				duration.Hour = value
			case 6:
				duration.Minute = value
			case 7:
				duration.Second = value
			}
		}

		if field >= currentField {
			return Duration{}, fmt.Errorf("unexpected field `%s`: %w", token, ErrParsing)
		}
		field = currentField
	}

	if err := scanner.Err(); err != nil {
		return Duration{}, fmt.Errorf("scanner error `%s`: %w", err.Error(), ErrParsing)
	}

	if (Duration{}) == duration {
		return Duration{}, fmt.Errorf("invalid expression `%s`: %w", input, ErrParsing)
	}

	return duration, nil
}

// fields:
// P Y M D T H M S
func fieldIndex(s string, timePortion bool) int {
	switch timePortion {
	case false:
		switch s {
		case "P":
			return 0
		case "Y":
			return 1
		case "M":
			return 2
		case "D":
			return 3
		case "T":
			return 4
		}
	case true:
		switch s {
		case "H":
			return 5
		case "M":
			return 6
		case "S":
			return 7
		}
	}
	return -1
}

func isUnit(r rune) bool {
	switch r {
	case 'P', 'T', 'Y', 'M', 'D', 'H', 'S':
		return true
	default:
		return false
	}
}

func scanUnits(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isUnit(r) {
			return i + width, data[start : i+1], nil
		}
	}

	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	return start, nil, nil
}

// Strings implements fmt.Stringer
func (d Duration) String() string {
	var sb strings.Builder
	sb.WriteRune('P')

	if d.Year != 0 {
		sb.WriteString(strconv.FormatFloat(d.Year, 'f', -1, 64))
		sb.WriteRune('Y')
	}

	if d.Month != 0 {
		sb.WriteString(strconv.FormatFloat(d.Month, 'f', -1, 64))
		sb.WriteRune('M')
	}

	if d.Day != 0 {
		sb.WriteString(strconv.FormatFloat(d.Day, 'f', -1, 64))
		sb.WriteRune('D')
	}

	if d.Hour != 0 || d.Minute != 0 || d.Second != 0 {
		sb.WriteRune('T')
	}

	if d.Hour != 0 {
		sb.WriteString(strconv.FormatFloat(d.Hour, 'f', -1, 64))
		sb.WriteRune('H')
	}

	if d.Minute != 0 {
		sb.WriteString(strconv.FormatFloat(d.Minute, 'f', -1, 64))
		sb.WriteRune('M')
	}

	if d.Second != 0 {
		sb.WriteString(strconv.FormatFloat(d.Second, 'f', -1, 64))
		sb.WriteRune('S')
	}

	return sb.String()
}
