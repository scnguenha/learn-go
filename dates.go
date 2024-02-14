package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	past15DaysDate := time.Now().AddDate(0, 0, -15)
	fmt.Println(today)
	fmt.Println(past15DaysDate)
}
