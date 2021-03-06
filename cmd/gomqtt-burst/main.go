package main

import (
	"flag"
	"time"

	"github.com/256dpi/gomqtt/client"
	"github.com/256dpi/gomqtt/packet"
)

var urlString = flag.String("url", "tcp://0.0.0.0:1883", "broker url")
var topic = flag.String("topic", "burst", "the used topic")
var qos = flag.Uint("qos", 0, "the qos level")
var amount = flag.Int("n", 100, "the amount of messages")

func main() {
	flag.Parse()

	cl := client.New()

	cl.Callback = func(msg *packet.Message, err error) error {
		if err != nil {
			panic(err)
		}

		return nil
	}

	cf, err := cl.Connect(client.NewConfig(*urlString))
	if err != nil {
		panic(err)
	}

	err = cf.Wait(10 * time.Second)
	if err != nil {
		panic(err)
	}

	var futures []client.GenericFuture

	for i := 0; i < *amount; i++ {
		pf, err := cl.Publish(*topic, []byte(*topic), uint8(*qos), false)
		if err != nil {
			panic(err)
		}

		futures = append(futures, pf)
	}

	for _, pf := range futures {
		err = pf.Wait(10 * time.Second)
		if err != nil {
			panic(err)
		}
	}

	err = cl.Disconnect()
	if err != nil {
		panic(err)
	}
}
