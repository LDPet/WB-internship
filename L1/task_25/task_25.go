package main

import (
	"time"
)

func main() {
	t := time.Now()
	sleep(3 * time.Second)
	println(time.Since(t).String())
}

// sleep является активной блокировкой (горутина не блокируется, т.е. продолжает работать),
// в отличии от time.Sleep где вызываетсы gopark
func sleep(d time.Duration) {
	t := time.Now()
	for {
		// курутимся в цикле пока не пройдёт заданное время
		if time.Since(t) > d {
			return
		}
	}
}
