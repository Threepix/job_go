package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

const (
	hash_1 = "1115dd800feaacefdf481f1f9070374a2a81e27880f187396db67958b207cbad"
	hash_2 = "3a7bd3e2360a3d29eea436fcfb7e44c735d117c42d1c1835420b6b9942dd4f1b"
	hash_3 = "74e1bb62f8dabb8125a58852b63bdf6eaef667cb56ac7f7cdba6d7305c50a22f"
)

func randInt(min int, max int) int {

	return min + rand.Intn(max-min)
}

func randomString(len int) string {

	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(97, 123))
	}

	return string(bytes)
}

func hash1() string {
	start := time.Now()
	otw := ""
	for true {
		fan := randomString(5)
		//fmt.Println("1 " + fan)
		res := fmt.Sprintf("%x", sha256.Sum256([]byte(fan)))
		if res == hash_1 {
			otw += fan
			fmt.Println("1 is " + otw)
			break
		}
	}
	duration := time.Since(start)
	fmt.Println(duration)
	return otw
}

func main() {

	go hash1()

	go func() {
		start := time.Now()
		otw := ""
		for true {
			fan := randomString(5)
			//fmt.Println("1 " + fan)
			res := fmt.Sprintf("%x", sha256.Sum256([]byte(fan)))
			if res == hash_2 {
				otw += fan
				fmt.Println("2 is " + otw)
				break
			}
		}
		duration := time.Since(start)
		fmt.Println(duration)
	}()

	go func() {
		start := time.Now()
		otw := ""
		for true {
			fan := randomString(5)
			//fmt.Println("2 " + fan)
			res := fmt.Sprintf("%x", sha256.Sum256([]byte(fan)))
			if res == hash_3 {
				otw += fan
				fmt.Println("3 is " + otw)
				break
			}
		}
		duration := time.Since(start)
		fmt.Println(duration)
	}()

	time.Sleep(500 * time.Second)
}
