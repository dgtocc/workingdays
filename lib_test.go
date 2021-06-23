package workingdays

import (
	"log"
	"testing"
	"time"
)

func TestAfter(t *testing.T) {
	InitStr([]string{
		"23/06/21",
		"25/06/21",
		"27/06/21",
		"29/06/21",
		"06/07/21",
		"12/07/21",
		"08/07/21",
		"15/07/21",
	}, "02/01/06")
	for i := 0; i < 16; i++ {
		//i:=7
		log.Printf("===========")
		log.Printf("%d => %d %s", i, After(time.Now(), i), End(time.Now(), i).Format("02/01/06"))
	}
}
