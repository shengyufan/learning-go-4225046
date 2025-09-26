package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Dates and times")

	t := time.Date(1996, time.February, 25, 0, 0, 0, 0, time.UTC)
	fmt.Printf("This specified time is %s\n", t)

	now := time.Now()
	fmt.Printf("Current time is %s\n", now)
	fmt.Printf("Object type is %T\n", now)

	// 按照一定格式进行输出
	fmt.Printf("Now is %s\n", now.Format(time.ANSIC))
	// AddDate 直接进行日期的计算
	tomorrow := time.Now().AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %s\n", tomorrow.Format(time.ANSIC))

	// 自定义输出格式
	format := "Monday 2006-01-02"
	fmt.Printf("New format: %s\n", now.Format(format))
}
