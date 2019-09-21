package util

func ChanIsClosed(ch <-chan interface{}) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

//==================================================11
type Semaphore chan struct{}

func MakeSemaphore(n int) Semaphore {
	return make(Semaphore, n)
}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}

//==================================================11
