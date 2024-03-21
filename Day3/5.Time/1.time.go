package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time handing...")
	//Current

	t := time.Now()
	fmt.Println(t)

	// constructing time

	t = time.Date(2024, 3, 20, 14, 35, 0, 0, &time.Location{})

	fmt.Println(t)

	// formating time to string and string time
	// year/month/day hh:min:second

	layout := "2006/01/02 15:04:05"

	str := t.Format(layout)

	fmt.Println(str)

	t1, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t1)
	}

	// epoc linux time

	utime := t.Unix()
	fmt.Println(utime)
	t2 := time.Unix(utime, 0)
	fmt.Println(t2)

	// time addition and sub traction

	t3 := t.Add(10 * time.Minute)

	fmt.Println(t3)
}
