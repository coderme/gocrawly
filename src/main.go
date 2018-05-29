package main

import (
	"fmt"
	"time"
)

func main() {
	fetch(urls...)

	if *revisitEvery == 0 {
		return
	}

	tick := time.Tick(*revisitEvery)
	fmt.Printf("Revisiting every %s\n%s",
		*revisitEvery,
		"=========================\n",
	)

	for {
		select {
		case <-tick:
			go fetch(urls...)
		}
	}

}
