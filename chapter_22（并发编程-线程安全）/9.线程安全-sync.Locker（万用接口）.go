package main

import "sync"

func main() {
	var locker sync.Locker
	m := new(sync.Mutex)
	rw := new(sync.RWMutex)

	locker = m
	locker.Lock()   //互斥锁的加锁
	locker.Unlock() //互斥锁的解锁

	locker = rw
	locker.Lock()   //读写锁的加锁
	locker.Unlock() //读写锁的解锁
}
