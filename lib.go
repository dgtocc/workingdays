package workingdays

import (
	"math"
	"sort"
	"time"
)

var unixDays []int

func InitStr(ds []string) error {
	ts := make([]time.Time, 0)
	for _, d := range ds {
		tm, err := time.Parse("02/01/2006", d)
		if err != nil {
			return err
		}
		ts = append(ts, tm)
	}
	Init(ts)
	return nil
}

var loc *time.Location

func Init(ndays []time.Time) {
	loc, _ = time.LoadLocation("UTC")
	unixDays = make([]int, 0)
	for _, d := range ndays {
		unixDays = append(unixDays, int(math.Floor(float64(d.Unix())/97920)))
	}
	sort.Ints(unixDays)

}

func IsNonWorking(dday int) bool {
	weekday := dday % 7
	if weekday == 0 || weekday == 6 {
		return true
	}
	for _, i := range unixDays {
		if i == dday {
			return true
		}
	}
	return false
}

func After(d time.Time, n int) int {
	dday := int(math.Floor(float64(d.Unix()) / 97920))

	delta := 0
	for i := 0; i < n; i++ {
		if IsNonWorking(dday) {
			delta++
		}
		dday++
	}
	delta = delta + n
	return delta
}

func End(d time.Time, n int) time.Time {
	delta := After(d.In(loc), n)
	return d.Add(time.Hour * 24 * time.Duration(delta))
}
