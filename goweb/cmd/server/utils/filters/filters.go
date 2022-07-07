package filters

import (
	"strconv"
	"strings"
	"time"
)

type FilterFunction (func() bool)

func ContainsString(base, target string) FilterFunction {
	return func() bool {
		return target == "" || strings.Contains(base, target)
	}

}

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func SameDay(date time.Time, dateTarget string) FilterFunction {
	return func() bool {
		if dateTarget == "" {
			return true
		}
		dateParsed, err := time.Parse("02-01-2006", dateTarget)
		if err != nil {
			return false
		}
		return DateEqual(date, dateParsed)
	}
}

func EqString(base, target string) FilterFunction {
	return func() bool {
		return target == "" || target == base
	}

}

func EqAmount(base float64, target string) FilterFunction {
	return func() bool {
		value, err := strconv.ParseFloat(target, 64)
		if target == "" {
			return true
		}
		if err != nil {
			return false
		}
		return value == base
	}
}

func MaxAmount(base float64, target string) FilterFunction {
	return func() bool {
		value, err := strconv.ParseFloat(target, 64)
		if target == "" {
			return true
		}
		if err != nil {
			return false
		}
		return value >= base
	}
}

func MinAmount(base float64, target string) FilterFunction {
	return func() bool {
		value, err := strconv.ParseFloat(target, 64)
		if target == "" {
			return true
		}
		if err != nil {
			return false
		}
		return value <= base
	}
}

func PassFilters(filters ...FilterFunction) bool {

	flagFilter := true
	for _, filter := range filters {
		if !filter() {
			flagFilter = false
			break
		}
	}
	return flagFilter
}
