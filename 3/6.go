package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RC(min, max int) int {
	return min + rand.Intn(max-min)
}

func pop(list []int) []int {
	if len(list) > 0 {
		list = list[:len(list)-1]
	}
	return list
}

func post(m sync.Mutex, ch []int, quit chan string, number int, flag int) {
	m.Lock()
	if flag == 0 {
		fmt.Println("POST" + fmt.Sprintf("%d", number) + "started")
	}
	for i := 0; i < 1; i++ {
		if len(quit) > 0 {
			time.Sleep(1 * time.Second)
			m.Unlock()
			break
		}
		t := RC(1, 100)
		if len(ch) >= 100 {
			time.Sleep(1 * time.Second)
			m.Unlock()
			continue
		}
		if len(ch) <= 80 {
			time.Sleep(1 * time.Second)
			ch = append(ch, t)
			fmt.Println("POST" + fmt.Sprintf("%d", number) + "gave" + fmt.Sprintf("%d", t))
			m.Unlock()
		}
	}
}
func consumer(m sync.Mutex, ch []int, quit chan string, number int, flag int) {
	m.Lock()
	if flag == 0 {
		fmt.Println("CONSUMER" + fmt.Sprintf("%d", number) + "started")
	}
	for i := 0; i < 1; i++ {
		if len(ch) == 0 && len(quit) > 0 {
			time.Sleep(2 * time.Second)
			m.Unlock()
			break
		}
		if len(ch) > 1 && len(quit) == 0 {
			fmt.Println("CONSUMER " + fmt.Sprintf("%d", number) + " get" + fmt.Sprintf("%d", ch[len(ch)-1]))
			pop(ch)
			time.Sleep(2 * time.Second)
			m.Unlock()
		}
		if len(ch) == 0 && len(quit) == 0 {
			time.Sleep(2 * time.Second)
			m.Unlock()
			continue
		}
	}
}

func main() {
	var ch []int
	quit := make(chan string)
	flag := 0
	var m sync.Mutex

	for {
		for i := 0; i < 5; i++ {
			if i == 4 {
				go post(m, ch, quit, i+1, flag)
			}
			go post(m, ch, quit, i+1, flag)
			go consumer(m, ch, quit, i+1, flag)
			flag++
		}
	}
}
