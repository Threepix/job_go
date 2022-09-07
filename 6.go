package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math/rand"
	"os"
)

func Random(min int, max int) int {

	return min + rand.Intn(max)
}

func proizv1(ch chan int) {
	fmt.Println("post 1 started")

	fmt.Println("Press q to quit")
	for true {
		if len(ch) >= 100 {
			fmt.Println("Continuing loop")
			continue
		}
		if len(ch) <= 80 {
			t := Random(1, 100)
			ch <- t
			//fmt.Println("1 chanal make " + string(<-ch))
		}
	}
}
func proizv2(ch chan int) {
	fmt.Println("post 2 started")

	fmt.Println("Press q to quit")
	for true {
		if len(ch) >= 100 {
			fmt.Println("Continuing loop")
			continue
		}
		if len(ch) <= 80 {
			t := Random(1, 100)
			ch <- t
			//fmt.Println("2 chanal make " + string(<-ch))
		}
	}
}
func proizv3(ch chan int, dead chan bool) {
	fmt.Println("post3 started")
	fmt.Println("Press q to quit")
	t := Random(1, 100)
	for true {
		if len(ch) >= 100 {
			fmt.Println("Continuing loop")
			continue
		}
		select {
		case ch <- t:
			if len(ch) <= 80 {
				ch <- t
				//fmt.Println("3 chanal make " + string(<-ch))
			}
		case <-dead:
			close(ch)
			return
		}
	}
}

func pocup1(ch1 chan int) {
	fmt.Println("byer 1 started")
	for true {
		if len(ch1) > 0 {
			fmt.Println(<-ch1)
			<-ch1
		}
		continue
	}
}
func pocup2(ch1 chan int) {
	fmt.Println("byer 1 started")
	for true {
		if len(ch1) > 0 {
			fmt.Println(<-ch1)
			<-ch1
		}
		continue
	}
}

func main() {
	chan1 := make(chan int, 200)
	quit := make(chan bool)
	fmt.Println("chanal was created")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	/*fmt.Println("Press q to start")
	fmt.Println("Press q to quit")*/

	_, key, err := keyboard.GetKey()
	fmt.Println("Press q to start")
	if err != nil {
		panic(err)
	}

	go func() {
		fmt.Println("post1 started")
		for true {
			t := Random(1, 100)
			select {
			case chan1 <- t:
				if len(chan1) >= 100 {
					//fmt.Println("Continuing loop")
					continue
				}
				if len(chan1) <= 80 {
					chan1 <- t
					//fmt.Println("3 chanal make " + string(<-ch))
				}
			case <-quit:
				close(chan1)
				return
			}
		}
	}()
	go func() {
		fmt.Println("post2 started")
		for true {
			t := Random(1, 100)
			select {
			case chan1 <- t:
				if len(chan1) >= 100 {
					//fmt.Println("Continuing loop")
					continue
				}
				if len(chan1) <= 80 {
					chan1 <- t
					//fmt.Println("3 chanal make " + string(<-ch))
				}
			case <-quit:
				close(chan1)
				return
			}
		}
	}()
	go func() {
		fmt.Println("post3 started")
		for true {
			t := Random(1, 100)
			select {
			case chan1 <- t:
				if len(chan1) >= 100 {
					//fmt.Println("Continuing loop")
					continue
				}
				if len(chan1) <= 80 {
					chan1 <- t
					//fmt.Println("3 chanal make " + string(<-ch))
				}
			case <-quit:
				close(chan1)
				return
			}
		}
	}()

	//pokupatels

	go pocup1(chan1)

	go pocup2(chan1)

	defer func() {
		_ = keyboard.Close()
	}()

	go func() {
		for {
			if key == keyboard.KeyCtrlQ {
				fmt.Println("в потоке щас" + string(len(chan1)))
				go func() {
					quit <- true
				}()
				os.Exit(0)
			}
		}
	}()
}
