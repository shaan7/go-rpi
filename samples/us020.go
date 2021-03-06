package main

import (
	"log"
	"time"

	"github.com/kid0m4n/go-rpi/sensor/us020"
)

func main() {
	rf := us020.New(10, 9, nil)
	defer rf.Close()

	for {
		distance, err := rf.Distance()
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Distance is %v", distance)

		time.Sleep(500 * time.Millisecond)
	}
}
