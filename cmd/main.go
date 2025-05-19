package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/haquenafeem/performer"
)

func main() {
	p := performer.NewSimplePerformer()
	p.Perform(func(ctx *performer.Context) error {
		ctx.Set("a", []string{"b"})
		go func() {
			time.Sleep(time.Second * 5)
			fmt.Println("hello")
		}()

		return errors.New("step 1 err")
	}).Perform(func(ctx *performer.Context) error {
		err := ctx.Err()
		if err != nil {
			fmt.Println(err)
		}

		val, ok := ctx.Get("a")
		if !ok {
			fmt.Println("value not found")
		}
		fmt.Println(val)

		return errors.New("step 2")
	})
}
