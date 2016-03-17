package conc

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var initialString string
var finalString string
var stringLength int

func addToFinalStack(letterChannel chan string, wg *sync.WaitGroup) {
	letter := <-letterChannel
	finalString += letter
	wg.Done()
}

func capitalize(letterChannel chan string, currentLetter string,
	wg *sync.WaitGroup) {
	thisLetter := strings.ToUpper(currentLetter)
	wg.Done()
	letterChannel <- thisLetter
}

func StartChan() {
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup

	initialString = "Four score and seven years"
	initialBytes := []byte(initialString)
	var letterChannel chan string = make(chan string)
	stringLength = len(initialBytes)

	for i := 0; i < stringLength; i++ {
		wg.Add(2)
		go capitalize(letterChannel, string(initialBytes[i]), &wg)
		go addToFinalStack(letterChannel, &wg)
		wg.Wait()
	}
	fmt.Println(finalString)
}
