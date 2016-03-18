package conc

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

//A buffered channel is synchronous.
//Go’s channel system is based on Communicating Sequential Processes (CSP),
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
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup

	initialString = "When we first brainstormed Swarm Week"
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
