package myfin

type (
	Weekday  int
	Weekdays map[string]Weekday
)

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w *Weekdays) Has(day string) bool {
	t := *w

	_, ok := t[day]

	return ok
}
