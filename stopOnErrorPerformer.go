package performer

type stopOnErrorPerformer struct {
	ctx *Context
}

func (p *stopOnErrorPerformer) Perform(fn HandlerFunc) Performer {
	if p.ctx.Err() != nil {
		return p
	}

	err := fn(p.ctx)
	p.ctx.err = err

	return p
}

func NewStopOnErrorPerformer() Performer {
	return &stopOnErrorPerformer{
		ctx: &Context{
			dataMap: make(map[string]any),
			err:     nil,
		},
	}
}
