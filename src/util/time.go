package util

import (
	"fmt"
	"time"
)

func FormatTime() {
	// get current timestamp
	currentTime := time.Now().Local()

	//print time
	fmt.Println(currentTime)

	newFormat := currentTime.Format("2006-01-02 15:04:05.000")
	fmt.Println(newFormat)

	fmt.Println("milliseconds:", time.Now().UnixNano()/int64(time.Millisecond))

	//TODO Changing time layout(form)
}

func StartCac() {
	t1 := time.Now() // get current time
	for i := 0; i < 1000; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
