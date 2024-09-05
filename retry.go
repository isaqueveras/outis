package outis

// retry allowes a given method to be retried x amount of times.
type retry struct{ amount, retries int8 }

// Go returns the settings for using the Attempt method
func (ctx *Context) Retry(retries int8) *retry {
	return &retry{retries: retries}
}

// The attempt attempts the given method for a given amount of retries
// If the method still fails after the set limit is a error returned.
func (r *retry) Attempt(method func() error) (err error) {
	if err = method(); err == nil {
		return nil
	}

	if r.retries >= r.amount {
		return err
	}

	r.retries++
	return r.Attempt(method)
}
