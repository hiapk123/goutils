package atomicinteger

import (
	"sync"
	"sync/atomic"
)

type AtomicInteger struct {
	value int32
	lock  sync.Mutex
}

func (atom *AtomicInteger) init() *AtomicInteger {
	atom.value = 0
	return atom
}

func NewAtomicInteger() *AtomicInteger {
	return new(AtomicInteger).init()
}

func (atom *AtomicInteger) Get() int32 {
	return atom.value
}

func (atom *AtomicInteger) Set(newValue int32) {
	atom.value = newValue
}

func (atom *AtomicInteger) GetAndSet(newValue int32) int32 {
	oldValue := atom.value
	atomic.AddInt32(&atom.value, newValue)
	return oldValue
}

func (atom *AtomicInteger) CompareAndSet(except, update int32) {
	atomic.CompareAndSwapInt32(&atom.value, except, update)
}

func (atom *AtomicInteger) GetAndIncrement() int32 {
	atom.lock.Lock()
	defer atom.lock.Unlock()
	value := atomic.LoadInt32(&atom.value)
	atomic.StoreInt32(&atom.value, value+1)
	return value
}

func (atom *AtomicInteger) GetAndDecrement() int32 {
	atom.lock.Lock()
	defer atom.lock.Unlock()
	value := atomic.LoadInt32(&atom.value)
	atomic.StoreInt32(&atom.value, value-1)
	return value
}

func (atom *AtomicInteger) IncrementAndGet() int32 {
	atom.lock.Lock()
	defer atom.lock.Unlock()
	atomic.StoreInt32(&atom.value, atom.value+1)
	return atomic.LoadInt32(&atom.value)
}

func (atom *AtomicInteger) DecrementAndGet() int32 {
	atom.lock.Lock()
	defer atom.lock.Unlock()
	atomic.StoreInt32(&atom.value, atom.value-1)
	return atomic.LoadInt32(&atom.value)
}
