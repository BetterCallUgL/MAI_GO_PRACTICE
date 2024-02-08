package rwmutex

// A RWMutex is a reader/writer mutual exclusion lock.
// The lock can be held by an arbitrary number of readers or a single writer.
// The zero value for a RWMutex is an unlocked mutex.
//
// If a goroutine holds a RWMutex for reading and another goroutine might
// call Lock, no goroutine should expect to be able to acquire a read lock
// until the initial read lock is released. In particular, this prohibits
// recursive read locking. This is to ensure that the lock eventually becomes
// available; a blocked Lock call excludes new readers from acquiring the
// lock.
type RWMutex struct {
	wr  chan struct{}
	r   chan struct{}
	cnt int
}

// New creates *RWMutex.
func New() *RWMutex {
	var rw RWMutex
	rw.wr = make(chan struct{}, 1)
	rw.r = make(chan struct{}, 1)
	rw.r <- struct{}{}
	rw.wr <- struct{}{}
	rw.cnt = 0
	return &rw
}

// RLock locks rw for reading.
//
// It should not be used for recursive read locking; a blocked Lock
// call excludes new readers from acquiring the lock. See the
// documentation on the RWMutex type.
func (rw *RWMutex) RLock() {
	<-rw.r
	if rw.cnt == 0 {
		<-rw.wr
	}
	rw.cnt++
	rw.r <- struct{}{}
}

// RUnlock undoes a single RLock call;
// it does not affect other simultaneous readers.
// It is a run-time error if rw is not locked for reading
// on entry to RUnlock.
func (rw *RWMutex) RUnlock() {
	<-rw.r // ЧИТАЛЕТЬ пытается 
	rw.cnt--
	if rw.cnt == 0 {
		rw.wr <- struct{}{}
	}
	rw.r <- struct{}{}
}

// Lock locks rw for writing.
// If the lock is already locked for reading or writing,
// Lock blocks until the lock is available.
func (rw *RWMutex) Lock() {
	for {
		<-rw.r
		if rw.cnt != 0 {
			rw.r <- struct{}{}
		} else {
			break
		}
	}
	<-rw.wr
}

// Unlock unlocks rw for writing. It is a run-time error if rw is
// not locked for writing on entry to Unlock.
//
// As with Mutexes, a locked RWMutex is not associated with a particular
// goroutine. One goroutine may RLock (Lock) a RWMutex and then
// arrange for another goroutine to RUnlock (Unlock) it.
func (rw *RWMutex) Unlock() {
	rw.wr <- struct{}{}
	rw.r <- struct{}{}
}
