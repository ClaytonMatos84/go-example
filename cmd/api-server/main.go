package main

import (
	"fmt"
	"runtime"
	"time"
)

func count(withGoRoutine bool) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, withGoRoutine)
		time.Sleep(time.Second)
	}
}

func testGoRoutine() {
	count(false)
	go count(true)
	fmt.Println("Hello 1")
	fmt.Println("Hello 2")

	time.Sleep(time.Second * 10)
	fmt.Println("end")
}

func testRuntime() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Start")

	go func() {
		for {
			testGoRoutine()
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("end")
}

func testChannel()  {
	queue := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			queue <- i
			i++
		}
	}()

	for x := range queue {
		fmt.Println(x)
	}
}

func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker: ", workerId, " Msg: ", res)
		time.Sleep(time.Second)
	}
}

func testWorker()  {
	msg := make(chan int)
	go worker(1 , msg)
	go worker(2 , msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}
}

func main() {
	// testGoRoutine()
	// testRuntime()
	// testChannel()
	testWorker()
}
