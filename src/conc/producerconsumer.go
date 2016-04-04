package conc

import (
	"fmt"
)

var connection = make(chan bool)
var did = make(chan bool)
var seq int = 0

func doInProducer() {
	for i := 0; i < 10; i++ {
		connection <- true
	}
	did <- true
}

func doInConsumer() {
	for {
		communication := <-connection
		seq++
		fmt.Println("Communication from producer received!", communication, ", seq:", seq)
	}
}

func StartPC() {
	go doInProducer()
	go doInConsumer()
	<-did
	fmt.Println("All Done!")
}
