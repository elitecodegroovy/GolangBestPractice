// blocking.go
package conc

import (
	"fmt"
	"strconv"
	"time"
)

func thinkAboutKey(bC chan bool) {
	i := 0
	max := 5
	for {
		if i >= max {
			bC <- true
		}
		fmt.Println("Still Thinking")
		time.Sleep(1 * time.Millisecond)
		i++
	}
}

//Blocking method 1 – a listening, waiting channel
func blockSignal() {
	blockChannel := make(chan bool)
	go thinkAboutKey(blockChannel)

	//Despite the fact that all of our looping code is concurrent,
	// we’re waiting on a signal for
	//our blockChannel to continue linear execution.
	<-blockChannel
	fmt.Println("Blocking over!")
}

func doListener(fc chan func() string) {
	fc <- func() string {
		return "It has been sent!"
	}
}

//1 Creating a function channel
func blockFuncChannel() {
	listener := make(chan func() string)
	defer close(listener)
	go doListener(listener)

	select {
	case action := <-listener:
		message := action()
		fmt.Println("Message:", message, ", received!")
	}
}

//2 Using an interface channel
func blockingInterfaceChan() {
	mActor := make(chan MessageActor)
	//	defer close(mActor)
	for i := 0; i < 10; i++ {
		go sendMessage(mActor, i)
	}

	select {
	case message := <-mActor:
		fmt.Println(message.Send())
	}
	<-mActor
}

type MessageActor interface {
	Send() string
}

type Delivery struct {
	status string
}

func (d Delivery) Send() string {
	return d.status
}

func sendMessage(actor chan MessageActor, i int) {
	d := new(Delivery)
	d.status = "Did it with number " + strconv.FormatInt(int64(i), 10)
	actor <- d
}

func StartBlockingMode() {
	blockSignal()

	//func channel
	blockFuncChannel()

	//interface channel
	blockingInterfaceChan()
}
