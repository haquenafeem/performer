package performer

type Context struct {
	dataMap map[string]any
	err     error
}

func (ctx *Context) Get(key string) (any, bool) {
	value, ok := ctx.dataMap[key]
	if !ok {
		return nil, false
	}

	return value, true
}

func (ctx *Context) Set(key string, val any) {
	ctx.dataMap[key] = val
}

func (ctx *Context) Err() error {
	return ctx.err
}
