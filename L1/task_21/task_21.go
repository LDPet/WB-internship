package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// адптер решает проблему несовместимости используемого и предоставляемого интерфейса

// Target используемый интерфейс
type Target interface {
	Request() string
}

// Adaptee адапптируемая структура
type Adaptee struct {
}

// AdaptedOperation возвращает int а у Request string
func (adaptee *Adaptee) AdaptedOperation() int {
	return rand.Intn(100)
}

type Adapter struct {
	a Adaptee
}

// Request адаптиция интерфейса
func (a *Adapter) Request() string {
	return strconv.Itoa(a.a.AdaptedOperation())
}

func NewAdapter(adaptee Adaptee) Adapter {
	return Adapter{
		a: adaptee,
	}
}

func main() {
	adapter := NewAdapter(Adaptee{})

	fmt.Printf("%s", adapter.Request())
}
