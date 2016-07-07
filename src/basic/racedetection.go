package basic

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var transactionNo int

func transaction(amt int) bool {
	//mutex.Lock()
	approved := false
	if (balance - amt) < 0 {
		approved = false
	} else {
		approved = true
		balance = balance - amt
	}
	approvedText := "declined"
	if approved == true {
		approvedText = "approved"
	} else {
	}
	transactionNo = transactionNo + 1
	fmt.Println(transactionNo, "Transaction for $", amt, approvedText)
	fmt.Println("\tRemaining balance $", balance)
	//mutex.unlock()
	return approved
}

func DetectBalance() {
	rand.Seed(time.Now().Unix())

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	tranChan := make(chan bool)

	balance = 1000
	transactionNo = 0
	fmt.Println("Starting balance: $", balance)
	wg.Add(1)
	for i := 0; i < 100; i++ {
		go func(ii int, trChan chan (bool)) {
			transactionAmount := rand.Intn(25)
			transaction(transactionAmount)
			if ii == 99 {
				trChan <- true
			}
		}(i, tranChan)
	}
	go transaction(0)
	select {
	case <-tranChan:
		fmt.Println("Transactions finished")
		wg.Done()
	}
	wg.Wait()
	close(tranChan)
	fmt.Println("Final balance: $", balance)
}

func StartMuxBalance() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	balanceChan := make(chan int)
	tranChan := make(chan bool)
	balance = 1000

	fmt.Println("Starting balance: $", balance)
	wg.Add(1)
	for i := 0; i < 100; i++ {
		go func(ii int) {
			transactionAmount := rand.Intn(25)
			balanceChan <- transactionAmount
			if ii == 99 {
				fmt.Println("Should be quittin time")
				tranChan <- true
				close(balanceChan)
				wg.Done()
			}
		}(i)
	}
	go transaction(0)
	breakPoint := false
	for {
		if breakPoint == true {
			break
		}
		select {
		//Channels can be seen as the sanctioned go-to way of handling synchronized
		//dataUse Sync.Mutex().
		case amt := <-balanceChan:
			fmt.Println("Transaction for $", amt)
			if (balance - amt) < 0 {
				fmt.Println("Transaction failed!")
			} else {
				balance = balance - amt
				fmt.Println("Transaction succeeded")
			}
			fmt.Println("Balance now $", balance)
		case status := <-tranChan:
			if status == true {
				fmt.Println("Done")
				breakPoint = true
				close(tranChan)
			}
		}
	}
	wg.Wait()
	fmt.Println("Final balance: $", balance)
}
