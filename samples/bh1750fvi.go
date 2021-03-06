package main

import (
	"log"
	"time"

	"github.com/kid0m4n/go-rpi/i2c"
	"github.com/kid0m4n/go-rpi/sensor/bh1750fvi"
)

func main() {
	bus, err := i2c.NewBus(1)
	if err != nil {
		log.Panic(err)
	}

	sensor := bh1750fvi.New(bh1750fvi.High, bus)
	defer sensor.Close()

	for {
		lighting, err := sensor.Lighting()
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Lighting is %v lx", lighting)

		time.Sleep(500 * time.Millisecond)
	}
}
