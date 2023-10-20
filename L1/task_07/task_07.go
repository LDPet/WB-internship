package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// AsyncMapCacher структура для хранения мапы с мьютексом
type AsyncMapCacher[K comparable, V any] struct {
	hashMap map[K]V
	mutex   sync.RWMutex // для неблокируюзего чтения
}

func NewAsyncMapCacher[K comparable, V any]() *AsyncMapCacher[K, V] {
	return &AsyncMapCacher[K, V]{
		hashMap: make(map[K]V),
	}
}

func (a *AsyncMapCacher[K, V]) Set(key K, value V) {
	//захват мьютекса на запись
	a.mutex.Lock()
	//крит секция
	a.hashMap[key] = value
	//освобожденеи мьютекса
	a.mutex.Unlock()
}

func (a *AsyncMapCacher[K, V]) Get(key K) (V, bool) {
	//захват мьютекса на чтение
	// позволяет сразу нескольким горутинам читать, при условии отсутсвия записи в данный момент
	// также запрщает писать во время чтения
	a.mutex.RLock()
	//крит секция
	val, ok := a.hashMap[key]
	//освобожденеи мьютекса
	a.mutex.RUnlock()

	return val, ok
}

func (a *AsyncMapCacher[K, V]) Delete(key K) {
	//захват мьютекса на запись
	a.mutex.Lock()
	//крит секция
	delete(a.hashMap, key)
	//освобожденеи мьютекса
	a.mutex.Unlock()
}

func main() {
	m := NewAsyncMapCacher[int, int]()

	go func() {
		for {
			i := rand.Intn(1000)
			val := rand.Intn(1000)

			m.Set(i, val)
			fmt.Printf("Set \ti: %d \tval: %d\n", i, val)
		}
	}()

	go func() {
		for {
			i := rand.Intn(1000)
			val, ok := m.Get(i)

			if ok != true {
				fmt.Printf("No Get \ti: %d \t", i)
			} else {
				fmt.Printf("Get \ti: %d \tval: %d\n", i, val)
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
