package interruption

// Interruption implements an interruption semantic, similar to context.Context,
// but specialized for generic interruption (e.g. cancelling, pausing). See
// https://dave.cheney.net/2017/08/20/context-isnt-for-cancellation for
// reasoning on why context is not for cancellation.
type Interruption struct {
	ch chan struct{}
}

// New returns an initialized Interruption object.
func New() (*Interruption, func()) {
	c := Interruption{
		ch: make(chan struct{}),
	}
	return &c, c.interrupt
}

// Done checks if interruption was requested.
func (c *Interruption) Done() <-chan struct{} {
	return c.ch
}

// DoneNonBlock checks if interruption was requested, in a non-blocking manner.
func (c *Interruption) DoneNonBlock() bool {
	select {
	case _, ok := <-c.ch:
		return !ok
	default:
		return false
	}
}

// interrupt sends an interruption request. It is private, so that the subjects
// of the interruption cannot call it and interrupts itself and everything else
// that is bound to it.
func (c *Interruption) interrupt() {
	close(c.ch)
}
