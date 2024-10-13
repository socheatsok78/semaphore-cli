package semaphore

type SemaphoreOptFn func(*Semaphore)

func NewWithOpts(opts ...SemaphoreOptFn) *Semaphore {
	s := &Semaphore{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
