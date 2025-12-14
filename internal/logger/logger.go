package logger

import (
	r "canopy/internal/service"
	"fmt"
	"time"
)

func Logman() {
	var cntSd, cntRich, cntDb int

	for {
		a := len(r.SliceForSd)
		if a != cntSd {
			fmt.Println("Count SliceForSD has changed: ", a)
			for i := cntSd; i < a; i++ {
				fmt.Println(r.SliceForSd[i].Id, r.SliceForSd[i].CreatedBy)
			}
			cntSd = a
		}

		b := len(r.SliceForRich)
		if b != cntRich {
			fmt.Println("Count SliceForRich has changed: ", b)
			for i := cntRich; i < b; i++ {
				fmt.Println(r.SliceForRich[i].Id, r.SliceForRich[i].CreatedBy)
			}
			cntRich = b
		}

		c := len(r.SliceForDb)
		if c != cntDb {
			fmt.Println("Count SliceForDb has changed: ", c)
			for i := cntDb; i < c; i++ {
				fmt.Println(r.SliceForDb[i].Id, r.SliceForDb[i].CreatedBy)
			}
			cntDb = c
		}

		time.Sleep(200 * time.Millisecond)
	}
}
