package performer

type asyncPerformer struct {
	ctx *Context
}

func (p *asyncPerformer) Perform(fn HandlerFunc) Performer {
	go func() {
		err := fn(p.ctx)
		p.ctx.Lock()
		p.ctx.err = err
		p.ctx.Unlock()
	}()

	return p
}

func NewAsyncPerformer() Performer {
	return &asyncPerformer{
		ctx: &Context{
			useLock: true,
			dataMap: make(map[string]any),
			err:     nil,
		},
	}
}
