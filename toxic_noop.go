package main

// The NoopToxic passes all data through without any toxic effects.
type NoopToxic struct{}

func (t *NoopToxic) Name() string {
	return "noop"
}

func (t *NoopToxic) IsEnabled() bool {
	return true
}

func (t *NoopToxic) Pipe(stub *ToxicStub) {
	for {
		select {
		case <-stub.interrupt:
			return
		case buf := <-stub.input:
			if buf == nil {
				stub.Close()
				return
			}
			stub.output <- buf
		}
	}
}