package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math/rand"
	"time"
)

func Random(min int, max int) int {

	return min + rand.Intn(max-min)
}

func proizv1(ch chan int) {
	fmt.Println("post 1 started")
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press q to quit")
	for true {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if len(ch) >= 100 {
			fmt.Println("Continuing loop")
			continue
		}
		if len(ch) <= 80 {
			t := Random(1, 100)
			ch <- t
			fmt.Println("1 chanal make " + string(<-ch))
		}
		if key == keyboard.KeyCtrlQ {
			break
		}
	}
}
func proizv2(ch chan int) {
	fmt.Println("post 2 started")
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press q to quit")
	for true {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if len(ch) >= 100 {
			fmt.Println("Continuing loop")
			continue
		}
		if len(ch) <= 80 {
			t := Random(1, 100)
			ch <- t
			fmt.Println("2 chanal make " + string(<-ch))
			time.Sleep(time.Second / 2)
		}
		if key == keyboard.KeyCtrlQ {
			break
		}
	}
}
func proizv3(ch chan int) {
	fmt.Println("post3 started")
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	fmt.Println("Press q to quit")
	for true {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if len(ch) >= 100 {
			fmt.Println("Continuing loop")
			continue
		}
		if len(ch) <= 80 {
			t := Random(1, 100)
			ch <- t
			fmt.Println("3 chanal make " + string(<-ch))
			time.Sleep(time.Second)
		}
		if key == keyboard.KeyCtrlQ {
			break
		}
	}
}

func pocup1(ch1, ch2, ch3 chan int) {
	fmt.Println("byer 1 started")
	for true {
		if len(ch1) > 0 || len(ch2) > 0 || len(ch3) > 0 {
			select {
			case <-ch1:
				fmt.Println(<-ch1)
				<-ch1
			case <-ch2:
				fmt.Println(<-ch2)
				<-ch2
			case <-ch3:
				fmt.Println(<-ch3)
				<-ch3
			}
		}
		continue
	}
}
func pocup2(ch1, ch2, ch3 chan int) {
	fmt.Println("byer 2 started")
	for true {
		if len(ch1) > 0 || len(ch2) > 0 || len(ch3) > 0 {
			select {
			case <-ch1:
				fmt.Println(<-ch1)
				<-ch1
			case <-ch2:
				fmt.Println(<-ch2)
				<-ch2
			case <-ch3:
				fmt.Println(<-ch3)
				<-ch3
			}
		}
		continue
	}
}

func main() {
	fmt.Println("вау я запустилось")
	chan1 := make(chan int, 200)
	chan2 := make(chan int, 200)
	chan3 := make(chan int, 200)
	fmt.Println("chanal was created")
	//proizvoditels

	go proizv1(chan1)

	go proizv2(chan2)

	go proizv3(chan3)

	//pokupatels

	go pocup1(chan1, chan2, chan3)

	go pocup2(chan1, chan2, chan3)

	time.Sleep(15 * time.Second)
}
