package main

import (
	"fmt"
	"github.com/bartier/producer-consumer/internal/buffer"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"sync"
)

func init() {
	logrus.SetLevel(logrus.InfoLevel)

	file, _ := os.OpenFile("producer_consumer.log", os.O_WRONLY|os.O_CREATE, 0755)
	logrus.SetOutput(file)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func consumer(bufferChan chan *buffer.Buffer) {
	logrus.Info("consumer()")

	for {
		b := <-bufferChan

		if !b.IsEmpty() {
			resource, _ := b.GetResource()

			logrus.Infof("main(): Consumer got resource '%d' from buffer", resource)
		} else {
			logrus.Warn("main(): Consumer can't get resource from empty buffer")
		}

	}
}

func producer(bufferChan chan *buffer.Buffer) {
	logrus.Info("producer()")

	sharedBuffer := buffer.Buffer{}

	for {
		randomResource := generateRandomInteger()

		if !sharedBuffer.IsFull() {
			sharedBuffer.SetResource(randomResource)

			logrus.Infof("main(): Producer inserted resource '%d' in buffer", randomResource)
			bufferChan <- &sharedBuffer

		}
	}
}

func generateRandomInteger() int {
	for {
		randomResource := rand.Intn(1000)

		if randomResource != 0 {
			return randomResource
		}
	}
}

var orchestrator sync.WaitGroup

func main() {
	fmt.Println("Producer Consumer Example in Go")

	var bufferChan = make(chan *buffer.Buffer)

	orchestrator.Add(2)

	go producer(bufferChan)
	go consumer(bufferChan)

	orchestrator.Wait()
}
