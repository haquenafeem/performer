package performer

type simplePerformer struct {
	ctx *Context
}

func (p *simplePerformer) Perform(fn HandlerFunc) Performer {
	err := fn(p.ctx)
	p.ctx.err = err

	return p
}

func NewSimplePerformer() Performer {
	return &simplePerformer{
		ctx: &Context{
			dataMap: make(map[string]any),
			err:     nil,
		},
	}
}
