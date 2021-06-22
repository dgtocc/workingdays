package workingdays

import (
	"log"
	"testing"
	"time"
)

func TestAfter(t *testing.T) {
	InitStr([]string{
		"23/06/2021",
		"25/06/2021",
		"27/06/2021",
	})
	for i := 0; i < 15; i++ {
		log.Printf("%d => %d %s", i, After(time.Now(), i), End(time.Now(), i).String())
	}
}
