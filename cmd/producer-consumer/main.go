package main

import (
	"fmt"
	"github.com/bartier/producer-consumer/internal/buffer"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"sync"
	"time"
)

func init() {
	logrus.SetLevel(logrus.InfoLevel)

	file, _ := os.OpenFile("producer_consumer.log", os.O_WRONLY|os.O_CREATE, 0755)
	logrus.SetOutput(file)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func consumer(bufferChan chan *buffer.Buffer, id int) {
	logrus.Info("consumer()")

	for {
		b := <-bufferChan

		if !b.IsEmpty() {
			resource, _ := b.GetResource()
			time.Sleep(time.Duration(rand.Int63n(10))*time.Millisecond)

			logrus.Infof("main(): Consumer ID %d get resource '%d'. Buffer now: %v", id, resource, b.Resources)
		} else {
			logrus.Warn("main(): Consumer can't get resource from empty buffer")
		}

	}
}

func producer(bufferChan chan *buffer.Buffer, id int, sharedBuffer *buffer.Buffer) {
	logrus.Info("producer()")

	for {
		randomResource := generateRandomInteger()
		time.Sleep(time.Duration(rand.Int63n(10))*time.Millisecond)

		if !sharedBuffer.IsFull() {
			sharedBuffer.SetResource(randomResource)

			logrus.Infof("main(): Producer ID %d put resource '%d'. Buffer now: %v", id, randomResource, sharedBuffer.Resources)
			bufferChan <- sharedBuffer

		} else {
			logrus.Infof("main(): Producer ID %d buffer full  '%d'. Buffer now: %v", id, randomResource, sharedBuffer.Resources)
		}
	}
}

func generateRandomInteger() int {
	for {
		randomResource := rand.Intn(1000)

		if randomResource >= 100  {
			return randomResource
		}
	}
}

var orchestrator sync.WaitGroup

func main() {
	fmt.Println("Producer Consumer Example in Go")

	var bufferChan = make(chan *buffer.Buffer)

	var buffer = new(buffer.Buffer)

	orchestrator.Add(2)

	go producer(bufferChan, 1, buffer)
	go producer(bufferChan, 2, buffer)
	go producer(bufferChan, 3, buffer)
	go producer(bufferChan, 4, buffer)
	go producer(bufferChan, 5, buffer)
	go producer(bufferChan, 6, buffer)
	go consumer(bufferChan, 1)

	orchestrator.Wait()
}
