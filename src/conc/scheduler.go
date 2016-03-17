package conc

import (
	"fmt"
	"io/ioutil"
	"runtime"
	//	"strconv"
	"time"
)

func ReadText(j *Job) {
	fileName := j.text + ".txt"
	fileContents := ""
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fileContents += j.text
		fmt.Println(j.text)
		j.i++
	}
	err := ioutil.WriteFile(fileName, []byte(fileContents), 0644)
	if err != nil {
		panic("Something went awry")
	}
}

func showNumber(num int) {
	//	tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	//	fmt.Println(num, "-time:", tstamp)
	fmt.Print("\t", num)
}

func StartReadText() {
	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3

	world := new(Job)
	world.text = "world"
	world.i = 0
	world.max = 5
	go ReadText(hello)
	go ReadText(world)

	cpus := runtime.GOMAXPROCS(0)
	fmt.Println(runtime.GOOS, runtime.NumCPU(), runtime.NumCgoCall(), runtime.NumGoroutine(), "list cpu number:", cpus)

	max := 8
	for i := 1; i < max; i++ {
		go showNumber(i)
	}
	//This is because there’s no guarantee as to how
	//many goroutines, if any, will complete before the end of the main() function
	runtime.Gosched()
	fmt.Println("exits!")
}
