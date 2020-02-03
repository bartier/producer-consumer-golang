package buffer

import (
	"errors"
	"github.com/sirupsen/logrus"
)

const (
	bufferSize = 5
)

type Buffer struct {
	Resources    [bufferSize]int
}

func (b *Buffer) GetSize() int {
	return bufferSize
}

func (b *Buffer) IsFull() bool {

	logrus.Debug("IsFull() method from Buffer")

	for i, v := range b.Resources {
		logrus.Debugf("IsFull(): index: %d, value: %d", i, v)

		if v == 0 {
			logrus.Debugf("IsFull(): Found an empty index. index: %d, value: %d", i, v)
			return false
		}
	}

	return true
}

func (b *Buffer) IsEmpty() bool {
	logrus.Debug("IsEmpty() method from Buffer")

	for i, v := range b.Resources {
		logrus.Debugf("IsEmpty(): index: %d, value: %d", i, v)

		if v != 0 {
			return false
		}
	}

	return true
}

func (b *Buffer) SetResource(n int) {
	logrus.Debug("SetResource() method from Buffer")

	for i, v := range b.Resources {
		logrus.Debugf("SetResource(): index: %d, value: %d", i, v)

		if v == 0  {
			logrus.Debugf("Found empty place to new resource '%d' ", n)
			b.Resources[i] = n
			break
		}
	}
}

func (b *Buffer) GetResource() (int, error)  {
	logrus.Debug("GetResource() method from Buffer")

	for i := len(b.Resources) - 1; i >= 0; i-- {
		value := b.Resources[i]
		logrus.Debugf("GetResource(): index: %d, value: %d", i, value)

		if value != 0 {
			b.Resources[i] = 0
			return value, nil
		}
	}

	return -1, errors.New("GetResource(): there is no resource to give to consumer")
}
