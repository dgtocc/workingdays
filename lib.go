package workingdays

import (
	"log"
	"sort"
	"time"
)

var unixDays []int
var ZERO time.Time = time.Unix(0, 0)
var loc *time.Location

func DayStr(d int) string {
	ddaytime := time.Unix(int64(d*24*60*60), 0).In(loc)
	return ddaytime.In(loc).String()

}

func DayInt(d time.Time) int {
	dday := int(d.Sub(ZERO).Hours() / 24)
	return dday

}

func InitStr(ds []string) error {
	loc, _ = time.LoadLocation("UTC")

	ts := make([]time.Time, 0)
	for _, d := range ds {
		tm, err := time.Parse("02/01/2006", d)
		if err != nil {
			return err
		}
		ts = append(ts, tm.In(loc))
	}
	Init(ts)
	return nil
}

func Init(ndays []time.Time) {
	unixDays = make([]int, 0)
	for _, d := range ndays {
		intd := DayInt(d)
		unixDays = append(unixDays, intd)
		log.Printf("Adding nonworking days: %s", DayStr(intd))
	}
	sort.Ints(unixDays)

}

func IsNonWorking(dday int) bool {
	weekday := dday % 7
	if weekday == 3 || weekday == 2 {
		log.Printf("	=> %s - %d - weekend", DayStr(dday), weekday)
		return true
	}
	for _, i := range unixDays {
		if i == dday {
			log.Printf("	=> %s - %s - marked non working", DayStr(i), DayStr(dday))
			return true
		}
	}
	log.Printf("	=> %s - Working Day", DayStr(dday))
	return false
}

func After(d time.Time, n int) int {

	dday := DayInt(d)
	delta := 0
	for i := 0; i < n; i++ {
		if IsNonWorking(dday + 1) {
			delta++
		}
		dday++
	}
	delta = delta + n
	return delta
}

func End(d time.Time, n int) time.Time {
	delta := After(d, n)
	return d.Add(time.Hour * 24 * time.Duration(delta)).In(loc)
}
