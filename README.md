# Duration

ISO 8601 duration parse

## Example

```go
duration, err := duration.Parse("P3Y6M4DT12H30M5S")
if err != nil {
    ...
}
fmt.Printf("%#v\n", data)
// Output: duration.Duration{Year:3, Month:6, Day:4, Hour:12, Minute:30, Second:5}
```

## Format

ISO 8601 duration format
ISO 8601 Durations are expressed using the following format, where (n) is replaced by the value for each of the date and time elements that follow the (n):

`P(n)Y(n)M(n)DT(n)H(n)M(n)S`

Where:

- P is the duration designator (referred to as "period"), and is always placed at the beginning of the duration.
- Y is the year designator that follows the value for the number of years.
- M is the month designator that follows the value for the number of months.
- W is the week designator that follows the value for the number of weeks.
- D is the day designator that follows the value for the number of days.
- T is the time designator that precedes the time components.
- H is the hour designator that follows the value for the number of hours.
- M is the minute designator that follows the value for the number of minutes.
- S is the second designator that follows the value for the number of seconds.

## License

GNU General Public License v3.0 or later

See [LICENSE](LICENSE) to see the full text.