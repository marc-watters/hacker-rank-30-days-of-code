package main

import "fmt"

func main() {
	var rDay, rMon, rYr int
	if _, err := fmt.Scan(&rDay, &rMon, &rYr); err != nil {
		panic(err)
	}

	var eDay, eMon, eYr int
	if _, err := fmt.Scan(&eDay, &eMon, &eYr); err != nil {
		panic(err)
	}

	var fine int
	if rYr > eYr {
		fine = 10000
	} else if rYr == eYr && rMon > eMon {
		fine = 500 * (rMon - eMon)
	} else if rYr == eYr && rMon == eMon && rDay > eDay {
		fine = 15 * (rDay - eDay)
	}

	fmt.Println(fine)
}
