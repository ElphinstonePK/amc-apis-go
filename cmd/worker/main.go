package main

import (
	"log"
	"math/rand"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const (
	address   = "15.184.170.196:7233"
	namespace = "default"
	taskQueue = "meezan"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	c, err := client.NewClient(client.Options{
		HostPort:  address,
		Namespace: namespace,
	})

	if err != nil {
		log.Fatalln("Unable To Create Temporal Client")
	}

	defer c.Close()

	w := worker.New(c, taskQueue, worker.Options{})
	register(w)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to Run Worker, Check Main", err)
	}
}
