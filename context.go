package performer

import "sync"

type Context struct {
	dataMap map[string]any
	err     error
	useLock bool
	sync.Mutex
}

func (ctx *Context) Get(key string) (any, bool) {
	if ctx.useLock {
		ctx.Lock()
		defer ctx.Unlock()
	}

	value, ok := ctx.dataMap[key]
	if !ok {
		return nil, false
	}

	return value, true
}

func (ctx *Context) Set(key string, val any) {
	if ctx.useLock {
		ctx.Lock()
		defer ctx.Unlock()
	}

	ctx.dataMap[key] = val
}

func (ctx *Context) Err() error {
	return ctx.err
}
