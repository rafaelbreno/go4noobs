#### Iota
- [Doc](https://golang.org/ref/spec#Iota)
- By _first sight_ looks like a counter inside const's scope
```
  const (
      a = iota // iota == 0
      b = iota // iota == 1
      c = iota // iota == 2
      d = iota // iota == 3
  )
```

##### Enum
- With Iota we can a ENUM-ish thing
```go
type day int

const (
	// Just some days of the week
	//
	SUNDAY day = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

func (d day) ToString() string {
	switch d {
	case SUNDAY:
		return "sunday"
	case MONDAY:
		return "monday"
	case TUESDAY:
		return "tuesday"
	case WEDNESDAY:
		return "wednesday"
	case THURSDAY:
		return "thursday"
	case FRIDAY:
		return "friday"
	case SATURDAY:
		return "saturday"
	default:
		return ""
	}
}

func main() {
	monday := MONDAY

	fmt.Println(monday.ToString())
}
```
