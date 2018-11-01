package atomiclong

import (
	"sync"
	"sync/atomic"
)

type AtomicLong struct {
	value int64
	lock  sync.Mutex
}

func (atom *AtomicLong) initInt64() *AtomicLong {
	atom.value = 0
	return atom
}

func NewAtomicLong() *AtomicLong {
	return new(AtomicLong).initInt64()
}

func (atom *AtomicLong) Get() int64 {
	return atom.value
}

func (atom *AtomicLong) Set(newValue int64) {
	atom.value = newValue
}

func (atom *AtomicLong) GetAndSet(newValue int64) int64 {
	oldValue := atom.value
	atomic.AddInt64(&atom.value, newValue)
	return oldValue
}

func (atom *AtomicLong) CompareAndSet(except, update int64) {
	atomic.CompareAndSwapInt64(&atom.value, except, update)
}

func (atom *AtomicLong) GetAndIncrement() int64 {
	atom.lock.Lock()
	defer atom.lock.Unlock()
	value := atomic.LoadInt64(&atom.value)
	atomic.StoreInt64(&atom.value, value+1)
	return value
}

func (atom *AtomicLong) GetAndDecrement() int64 {
	atom.lock.Lock()
	defer atom.lock.Unlock()
	value := atomic.LoadInt64(&atom.value)
	atomic.StoreInt64(&atom.value, value-1)
	return value
}

func (atom *AtomicLong) IncrementAndGet() int64 {
	atom.lock.Lock()
	defer atom.lock.Unlock()
	atomic.StoreInt64(&atom.value, atom.value+1)
	return atomic.LoadInt64(&atom.value)
}

func (atom *AtomicLong) DecrementAndGet() int64 {
	atom.lock.Lock()
	defer atom.lock.Unlock()
	atomic.StoreInt64(&atom.value, atom.value-1)
	return atomic.LoadInt64(&atom.value)
}
