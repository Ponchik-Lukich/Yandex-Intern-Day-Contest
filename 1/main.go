package main

import "fmt"

func main() {
	var days int
	var weekDay string
	var weekDayInt int
	counter := 0
	sCounter := 1
	fmt.Scan(&days, &weekDay)
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i := 0; i < len(week); i++ {
		if weekDay == week[i] {
			weekDayInt = i
			break
		} else {
			continue
		}
	}
	for i := 0; i < len(week); i++ {
		if i < weekDayInt {
			fmt.Print(".. ")
		} else {
			counter++
			fmt.Printf(".%d ", counter)
		}
	}
	fmt.Print("\n")
	for true {
		counter++
		if counter/10 == 0 {
			fmt.Printf(".%d ", counter)
		} else {
			fmt.Printf("%d ", counter)
		}
		if sCounter%7 == 0 {
			fmt.Print("\n")
		}
		sCounter++
		if counter == days {
			break
		}
	}
}
