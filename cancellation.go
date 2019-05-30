package cancellation

// Cancellation implements a cancellation type, similar to context.Context, but
// specialized for cancellation. See
// https://dave.cheney.net/2017/08/20/context-isnt-for-cancellation for
// reasoning on why context is not for cancellation.
type Cancellation struct {
	ch chan struct{}
}

// New returns an initialized Cancellation object.
func New() Cancellation {
	return Cancellation{
		ch: make(chan struct{}),
	}
}

// Done checks if cancellation was requested.
func (c *Cancellation) Done() <-chan struct{} {
	return c.ch
}

// DoneNonBlock checks if cancellation was requested, in a non-blocking manner.
func (c *Cancellation) DoneNonBlock() bool {
	select {
	case _, ok := <-c.ch:
		return !ok
	default:
		return false
	}
}

// Cancel sends a cancellation request.
func (c *Cancellation) Cancel() {
	close(c.ch)
}
