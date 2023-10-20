package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	b := NewBadCounter()
	m := NewMutexCounter()
	a := NewAtomicCounter()
	num := 10000

	println("Bad before", b.Val())
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			b.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	println("Bad after", b.Val())

	println("Mutex before", m.Val())
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			m.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	println("Mutex after", m.Val())

	println("Atomic before", a.Val())
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			a.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	println("Atomic after", a.Val())
}

type Counter interface {
	Inc()
	Val() uint64
}

type BadCounter struct {
	counter uint64
}

func NewBadCounter() BadCounter {
	return BadCounter{
		counter: 0,
	}
}

func (b *BadCounter) Inc() {
	b.counter++
}

func (b *BadCounter) Val() uint64 {
	return b.counter
}

type MutexCounter struct {
	counter uint64
	// для синхронизации, остальные ждут пока писатель не отпустит мьютекс
	mtx sync.Mutex
}

func NewMutexCounter() MutexCounter {
	return MutexCounter{
		counter: 0,
	}
}

func (m *MutexCounter) Inc() {
	m.mtx.Lock()
	m.counter++
	m.mtx.Unlock()
}

func (m *MutexCounter) Val() uint64 {
	m.mtx.Lock()
	val := m.counter
	m.mtx.Unlock()
	return val
}

type AtomicCounter struct {
	counter *uint64
}

func NewAtomicCounter() AtomicCounter {
	return AtomicCounter{
		counter: new(uint64),
	}
}

func (a *AtomicCounter) Inc() {
	// атомарная операция, т.е. не прерывается
	atomic.AddUint64(a.counter, 1)
}

func (a *AtomicCounter) Val() uint64 {
	return atomic.LoadUint64(a.counter)
}
