package channel

//in select, it represents a receive operation on a channel.
//A select statement will block the application
//until some information is sent along the channel
import (
	"fmt"
	"strings"
	"sync"
)

var initialString string
var initialBytes []byte
var stringLength int
var finalString string
var lettersProcessed int
var appStatus bool
var wg sync.WaitGroup

func getLetters(gQ chan string) {
	for i := range initialBytes {
		gQ <- string(initialBytes[i])
	}
}

func capitalizeLetters(gQ chan string, sQ chan string) {
	for {
		if lettersProcessed >= stringLength {
			appStatus = false
			break
		}
		select {
		case letter := <-gQ:
			capitalLetter := strings.ToUpper(letter)
			finalString += capitalLetter
			lettersProcessed++
		}
	}
}

func StartChannel() {
	appStatus = true

	getQueue := make(chan string)
	stackQueue := make(chan string)
	initialString := "When we first brainstormed Swarm Week"

	initialBytes = []byte(initialString)
	stringLength = len(initialString)
	lettersProcessed = 0

	go getLetters(getQueue)
	capitalizeLetters(getQueue, stackQueue)
	close(getQueue)
	close(stackQueue)
	//loop util the appStaus was set to false.
	for {
		if appStatus == false {
			fmt.Println("StartChannel Done.")
			fmt.Println(finalString)
			break
		}
	}
}
