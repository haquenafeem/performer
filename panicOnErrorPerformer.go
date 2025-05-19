package performer

type panicOnErrorPerformer struct {
	ctx *Context
}

func (p *panicOnErrorPerformer) Perform(fn HandlerFunc) Performer {
	if p.ctx.Err() != nil {
		panic(p.ctx.Err())
	}

	err := fn(p.ctx)
	p.ctx.err = err

	return p
}

func NewPanicOnErrorPerformer() Performer {
	return &panicOnErrorPerformer{
		ctx: &Context{
			dataMap: make(map[string]any),
			err:     nil,
		},
	}
}
