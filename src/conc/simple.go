package conc

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
}

func outputTextWithSync(j *Job, goGroup *sync.WaitGroup) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	goGroup.Done()
}

//Start the concurrency job
func StartSimple() {
	hello := new(Job)
	world := new(Job)

	hello.text = "Go Go"
	hello.i = 0
	hello.max = 3
	world.text = "gorutines"
	world.i = 0
	world.max = 5

	//Execute the job with concurrecy
	go outputText(hello)
	outputText(world)
}

func StartSyncSimple() {
	goGroup := new(sync.WaitGroup)
	fmt.Println("Starting")
	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 2
	world := new(Job)
	world.text = "world"
	world.i = 0
	world.max = 2

	go outputTextWithSync(hello, goGroup)
	go outputTextWithSync(world, goGroup)
	goGroup.Add(2)
	goGroup.Wait()
}
