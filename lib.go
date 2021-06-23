package workingdays

import (
	"log"
	"sort"
	"time"
)

var unixDays []int
var ZERO time.Time
var loc *time.Location

func DayStr(d int) string {
	ddaytime := time.Unix(int64(d*24*60*60), 0).In(loc)
	return ddaytime.String()

}

//Retorna o numero do dia a contar de ZERO, sendo ZERO 01/Jan/1970
func DayInt(d time.Time) int {
	d, _ = time.Parse("02/01/06", d.In(loc).Format("02/01/06"))
	dday := int(d.In(loc).Sub(ZERO).Hours() / 24)
	//log.Printf("Dayint: %s => %d", d.String(), dday)
	return dday
}

//Inicializa função com strings
func InitStr(ds []string, format string) error {
	loc, _ = time.LoadLocation("UTC")

	ts := make([]time.Time, 0)
	for _, d := range ds {
		tm, err := time.Parse(format, d)
		if err != nil {
			return err
		}
		ts = append(ts, tm)
	}
	Init(ts)
	return nil
}

//Inicializa função com Array de de time.Time
func Init(ndays []time.Time) {
	loc, _ = time.LoadLocation("UTC")
	ZERO = time.Unix(0, 0).In(loc)
	unixDays = make([]int, 0)
	for _, d := range ndays {
		intd := DayInt(d)
		unixDays = append(unixDays, intd)
		//log.Printf("Adding nonworking days: %s", DayStr(intd))
	}
	sort.Ints(unixDays)

}

//Verifica se o dia é util ou não
//Pode ser otimizada, se as consultas forem sempre crescentes através
//de um memento[GoF]
func IsNonWorking(dday int) bool {
	weekday := dday % 7
	if weekday == 3 || weekday == 2 {
		log.Printf("		=> %s(%d) - %d - weekend", DayStr(dday), dday, weekday)
		return true
	}
	for _, i := range unixDays {
		if i == dday {
			log.Printf("		=> %s - %s (%d)- marked non working", DayStr(i), DayStr(dday), dday)
			return true
		}
		if i > dday {
			log.Printf("		=> %s (%d) - Working Day", DayStr(dday), dday)
			return false
		}
	}
	log.Printf("		=> %s (%d) - Working Day", DayStr(dday), dday)
	return false
}

//Dado data e no dias uteis, retornar no dias corridos
func After(d time.Time, n int) int {

	dday := DayInt(d)
	delta := 0
	//if n==0{
	//	return 0
	//}
	i := 0
	for n > 0 {
		//log.Printf("==> Evaluating: %s", DayStr(dday+i))
		if !IsNonWorking(dday + i) {
			n--
		}
		i++
	}

	delta = delta + i
	return delta
}

//Dada data inicio e no dias uteis, calcula data Fim
func End(d time.Time, n int) time.Time {
	delta := After(d, n)
	return d.Add(time.Hour * 24 * time.Duration(delta))
}
