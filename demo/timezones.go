package demo

import "fmt"

type TZ int

const (
	HOUR TZ = 60 * 60
	UTC  TZ = 0 * HOUR
	EST  TZ = -5 * HOUR
	CST  TZ = -6 * HOUR
)

var timeZones = map[TZ]string{
	UTC: "Universal Greenwich time",
	EST: "Eastern Standard time",
	CST: "Central Standard time",
}

func (t TZ) String() string {
	if z, ok := timeZones[t]; ok {
		return z
	}
	return ""
}

func InitTimezones() {
	fmt.Println(UTC)
	fmt.Println(-6 * HOUR)
}
