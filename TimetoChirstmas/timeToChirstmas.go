package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)

	futureDate := time.Date(2021, time.December, 25, 00, 00, 00, 00, time.UTC)

	diff := futureDate.Sub(now)

	milliseconds := int(diff.Milliseconds())
	fmt.Printf("Milliseconds to Chirstmas: %d \n", milliseconds)

	second := int(diff.Seconds())
	fmt.Printf("Seconds to Chirstmas : %d\n", second)

	mins := int(diff.Minutes())
	fmt.Printf("Minutes to Chirstmas: %d \n", mins)

	hrs := int(diff.Hours())
	fmt.Printf("Hours to Chirstmas : %d \n", hrs)

	days := int(diff.Hours() / 24)
	fmt.Printf("Days to Chirstmas : %d \n", days)

}
