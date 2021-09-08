package log

import (
	"errors"
	"testing"

	"github.com/juan-chan/log/logfields"
)

func TestLogger(t *testing.T) {
	Info("Hello, World!")
	Named("greet").With(logfields.String("date", "2021-09-08")).Info("Hi, my name's Reed, what's your name?")
	Infof("Hello, %s!", "World")
	Warn("Hello, World!", logfields.String("name", "Joan"), logfields.Int("age", 18))
	Error("Hello error", logfields.Error(errors.New("nil pointer exception")))
}
